package cache

import (
	"context"
	"crudl/internal/domain"
	"errors"

	"github.com/redis/go-redis/v9"
)

func (c *Cache) Get(ctx context.Context, shortURL string) (*domain.ShortURLRequest, error) {
	res, err := c.Redis.Get(ctx, shortURL).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, &domain.NotFoundError{Message: "short_url not found"}
		}

		return nil, err
	}

	return &domain.ShortURLRequest{
		OriginalURL: res,
		ShortCode: shortURL,
	}, nil
}
