package usecase

import (
	"context"
	"crudl/internal/adapters/cache"
	"crudl/internal/adapters/postgres"
	"crudl/internal/domain"
)

type Postgres interface {
	CreateOrder(ctx context.Context, sub domain.Order) error
	GetOrder(ctx context.Context, user_id string) (domain.Order, error)
}

type Cache interface {
	Get(orderUID string) (domain.Order, bool)
	Add(order domain.Order)
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
