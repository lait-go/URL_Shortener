package cache

import (
	"crudl/internal/domain"
	"fmt"

	"golang.org/x/net/context"
)

func (c *Cache) Set(ctx context.Context, data *domain.ShortURL) error {
	exists, err := c.Redis.Exists(ctx, data.ShortCode).Result()
	if err != nil {
		return fmt.Errorf("error to check if key exists: %w", err)
	}

	fmt.Println(exists)

	if exists > 1{
		return &domain.KeyExistsError{}
	}

	if err := c.Redis.Set(ctx, data.ShortCode, data.OriginalURL, 0).Err(); err != nil {
		return fmt.Errorf("error to set key: %w", err)
	}

	return nil
}
