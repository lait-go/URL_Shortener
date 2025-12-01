package usecase

import (
	"context"
	"crudl/internal/adapters/cache"
	"crudl/internal/adapters/postgres"
	"crudl/internal/domain"
)

type Postgres interface {
	Set(ctx context.Context, data *domain.ShortURL) error
	Get(ctx context.Context, shortURL string) (*domain.ShortURL, error)
}

type Cache interface {
	Set(ctx context.Context, data *domain.ShortURL) error
	Get(ctx context.Context, shortURL string) (*domain.ShortURLRequest, error)
}

type Profile struct {
	Postgres Postgres
	Cache    Cache
}

func NewProfile(postgres *postgres.Pool, cache *cache.Cache) *Profile {
	return &Profile{
		Postgres: postgres,
		Cache:    cache,
	}
}
