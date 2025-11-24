package config

import (
	"crudl/internal/adapters/cache"
	"crudl/internal/adapters/postgres"
	"crudl/pkg/http_server"
	"crudl/pkg/logger"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Postgres postgres.Config
	Cache    cache.Config
	Http     http_server.Config
	Logger   logger.Config
}

func InitConfig() (*Config, error) {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("не удалось загрузить .env файл: %w", err)
	}

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("не удалось обработать переменные окружения: %w", err)
	}

	return &cfg, nil
}
