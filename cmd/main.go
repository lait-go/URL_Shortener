package main

import (
	"context"
	"errors"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	handlersAnalytics "github.com/avraam311/url-shortener/internal/api/handlers/analytics"
	handlersURL "github.com/avraam311/url-shortener/internal/api/handlers/url"
	"github.com/avraam311/url-shortener/internal/api/server"
	repositoryAnalytics "github.com/avraam311/url-shortener/internal/repository/analytics"
	repositoryURL "github.com/avraam311/url-shortener/internal/repository/url"
	serviceAnalytics "github.com/avraam311/url-shortener/internal/service/analytics"
	serviceURL "github.com/avraam311/url-shortener/internal/service/url"

	"github.com/wb-go/wbf/config"
	"github.com/wb-go/wbf/dbpg"
	"github.com/wb-go/wbf/redis"
	"github.com/wb-go/wbf/retry"
	"github.com/wb-go/wbf/zlog"

	"github.com/go-playground/validator/v10"
)

const (
	configLocal = "config/local.yaml"
	env         = ".env"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	zlog.Init()
	cfg := config.New()
	cfg.Load(configLocal, env, "")
	val := validator.New()

	opts := &dbpg.Options{
		MaxOpenConns:    cfg.GetInt("db.max_open_conns"),
		MaxIdleConns:    cfg.GetInt("db.max_idle_conns"),
		ConnMaxLifetime: cfg.GetDuration("db.conn_max_lifetime"),
	}
	slavesDNSs := []string{}
	masterDNS := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.GetString("DB_USER"), cfg.GetString("DB_PASSWORD"),
		cfg.GetString("DB_HOST"), cfg.GetString("PORT"),
		cfg.GetString("DB_NAME"), cfg.GetString("DB_SSL_MODE"),
	)
	db, err := dbpg.New(masterDNS, slavesDNSs, opts)
	if err != nil {
		zlog.Logger.Fatal().Err(err).Msg("failed to connect to database")
	}

	redis := redis.New(cfg.GetString("REDIS_ADDRESS"), cfg.GetString("REDIS_PASSWORD"), cfg.GetInt("REDIS_DATABASE"))
	if err = redis.Ping(ctx).Err(); err != nil {
		zlog.Logger.Fatal().Err(err).Msg("failed to connect to redis")
	}

	retryStrategy := retry.Strategy{
		Attempts: cfg.GetInt("retry.attempts"),
		Delay:    cfg.GetDuration("retry.delay"),
		Backoff:  cfg.GetFloat64("retry.backoff"),
	}
	repoURL := repositoryURL.NewRepository(db)
	repoAnalytics := repositoryAnalytics.NewRepository(db)
	srvcURL := serviceURL.NewService(repoURL, redis, retryStrategy)
	srvcAnalytics := serviceAnalytics.NewService(repoAnalytics)
	handURL := handlersURL.NewHandler(srvcURL, val)
	handAnalytics := handlersAnalytics.NewHandler(srvcAnalytics, val)

	router := server.NewRouter(handURL, handAnalytics, srvcAnalytics, cfg.GetString("api.gin_mode"))
	srv := server.NewServer(cfg.GetString("server.port"), router)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			zlog.Logger.Fatal().Err(err).Msg("failed to run server")
		}
	}()
	zlog.Logger.Info().Msg("server is running")

	<-ctx.Done()
	zlog.Logger.Info().Msg("shutdown signal received")

	shutdownCtx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	zlog.Logger.Info().Msg("shutting down server")
	if err := srv.Shutdown(shutdownCtx); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to shutdown server")
	}
	if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
		zlog.Logger.Info().Msg("timeout exceeded, forcing shutdown")
	}

	if err := db.Master.Close(); err != nil {
		zlog.Logger.Printf("failed to close master DB: %v", err)
	}
	for i, s := range db.Slaves {
		if err := s.Close(); err != nil {
			zlog.Logger.Printf("failed to close slave DB %d: %v", i, err)
		}
	}
}
