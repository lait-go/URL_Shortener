package cache

import (
	"crudl/internal/domain"
	"fmt"

	"golang.org/x/net/context"
)

func (c *Cache) Set(ctx context.Context, data domain.CreateShortURLRequest) error {
	exists, err := c.Redis.Exists(ctx, data.CustomAlias).Result()
	if err != nil {
		return fmt.Errorf("error to check if key exists: %w", err)
	}

	if exists != 1 {
		return fmt.Errorf(domain.ErrorExists) //TODO: реализовать нормальную возврат ошибки через domain
	}

	if err := c.Redis.Set(ctx, data.CustomAlias, data.OriginalURL, 0).Err(); err != nil {
		return fmt.Errorf("error to set key: %w", err)
	}

	return nil
}
