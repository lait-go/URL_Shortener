package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string `envconfig:"REDIS_HOST" required:"true"`
	Port     string `envconfig:"REDIS_PORT" required:"true"`
	Password string `envconfig:"REDIS_PASSWORD" required:"true"`
	DB       int    `envconfig:"REDIS_DB" required:"true"`
}

type Cache struct {
	Redis *redis.Client
}

func New(ctx context.Context, cfg Config) (*Cache, error) {
	cash := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if status := cash.Ping(ctx); status.Err() != nil {
		return &Cache{}, fmt.Errorf("error to connect redis")
	}

	return &Cache{Redis: cash}, nil
}

func (c *Cache) Close() error {
	return c.Redis.Close()
}
