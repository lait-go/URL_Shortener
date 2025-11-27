package usecase

import (
	"context"
	"crudl/internal/adapters/cache"
	"crudl/internal/adapters/postgres"
	"crudl/internal/domain"
)

type Postgres interface {
	Set(ctx context.Context, data domain.ShortURL) error
}

type Cache interface {
	Set(ctx context.Context, data domain.CreateShortURLRequest) error
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
