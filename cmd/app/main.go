package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"crudl/config"
	"crudl/internal/adapters/cache"
	"crudl/internal/adapters/postgres"
	"crudl/internal/adapters/postgres/migrations"
	"crudl/internal/controlers/http"
	"crudl/internal/usecase"
	"crudl/pkg/http_server"
	"crudl/pkg/logger"
)

func main() {
	ctx := context.Background()
	
	c, err := config.InitConfig()
	if err != nil {
		slog.Error("config init failed", "error", err)
		os.Exit(1)
	}

	logger.Init(c.Logger)
	log := logger.Default()
	
	log.InfoContext(ctx, "logger initialized")
	log.InfoContext(ctx, "starting migration...")

	if err = migrations.RunMigrate("internal/adapters/postgres/migrations", c.Postgres); err != nil {
		log.ErrorContext(ctx, "migration failed", err)
		os.Exit(1)
	}
	
	log.InfoContext(ctx, "migration completed")
	log.InfoContext(ctx, "starting app...")

	if err = AppRun(ctx, c); err != nil {
		log.ErrorContext(ctx, "application run failed", err)
		os.Exit(1)
	}
}

func AppRun(ctx context.Context, c *config.Config) error {
	log := logger.Default()
	
	post, err := postgres.New(ctx, c.Postgres)
	if err != nil {
		return err
	}
	log.InfoContext(ctx, "postgres initialized")

	cache, err := cache.New(ctx, c.Cache)
	if err != nil {
		return err
	}
	log.InfoContext(ctx, "cache initialized")

	profile := usecase.NewProfile(post, cache)
	log.InfoContext(ctx, "usecase layer initialized")

	router := http.Router(profile, c.Http)
	log.InfoContext(ctx, "router initialized")

	server := http_server.New(router, c.Http)
	
	serverErr := make(chan error, 1)
	go func() {
		log.InfoContext(ctx, "HTTP server starting on port: "+c.Http.Port)
		if err := server.Run(); err != nil {
			serverErr <- err
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	select {
	case <-sig:
		log.InfoContext(ctx, "received shutdown signal")
	case err := <-serverErr:
		log.ErrorContext(ctx, "server error", err)
		return err
	}

	if closeErr := post.Close(); closeErr != nil {
		log.ErrorContext(ctx, "failed to close postgres connection", closeErr)
	}

	if err := server.Close(c.Http.ShutdownTimeout); err != nil {
		log.ErrorContext(ctx, "failed to close server gracefully", err)
		return err
	}

	log.InfoContext(ctx, "application shutdown completed")
	return nil
}