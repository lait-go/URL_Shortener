package url

import (
	"context"

	"github.com/avraam311/url-shortener/internal/models/domain"

	"github.com/wb-go/wbf/retry"
)

//go:generate mockgen -source=service.go -destination=../../mocks/url/service.go
type RepositoryURL interface {
	SaveShortURL(context.Context, *domain.URL) (string, error)
	GetFullURL(context.Context, string) (string, error)
	CheckIfShortURLIsUnique(context.Context, string) (bool, error)
}

type cache interface {
	SetWithRetry(ctx context.Context, strategy retry.Strategy, key string, value interface{}) error
	GetWithRetry(ctx context.Context, strategy retry.Strategy, key string) (string, error)
}

type ServiceURL struct {
	repo          RepositoryURL
	cache         cache
	retryStrategy retry.Strategy
}

func NewService(repo RepositoryURL, cache cache, retryStrategy retry.Strategy) *ServiceURL {
	return &ServiceURL{
		repo:          repo,
		cache:         cache,
		retryStrategy: retryStrategy,
	}
}
