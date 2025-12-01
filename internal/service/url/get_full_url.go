package url

import (
	"context"
	"errors"
	"fmt"

	"github.com/wb-go/wbf/zlog"

	"github.com/go-redis/redis/v8"
)

func (s *ServiceURL) GetFullURL(ctx context.Context, shortURL string) (string, error) {
	fullURL, err := s.cache.GetWithRetry(ctx, s.retryStrategy, shortURL)
	if err != nil && !errors.Is(err, redis.Nil) {
		zlog.Logger.Warn().Err(err).Str("short_url", shortURL).Msg("failed to get full url by short url from cache")
	}

	if errors.Is(err, redis.Nil) {
		fullURL, err = s.repo.GetFullURL(ctx, shortURL)
		if err != nil {
			return "", fmt.Errorf("service/get_full_url.go - %w", err)
		}
		if fullURL[:7] != "http://" {
			fullURL = "http://" + fullURL
		}
		err = s.cache.SetWithRetry(ctx, s.retryStrategy, shortURL, fullURL)
		if err != nil {
			zlog.Logger.Warn().Err(err).Str("short_url", shortURL).Msg("failed to cache shortURL - fullURL")
		}
	}

	return fullURL, nil
}
