package analytics

import (
	"context"

	"github.com/avraam311/url-shortener/internal/models/db"
	"github.com/avraam311/url-shortener/internal/models/dto"
)

//go:generate mockgen -source=service.go -destination=../../mocks/analytics/service.go
type RepositoryAnalytics interface {
	GetAnalytics(context.Context, string) ([]*db.Analytics, error)
	SaveAnalytics(context.Context, *dto.Analytics) error
}

type ServiceAnalytics struct {
	repo RepositoryAnalytics
}

func NewService(repo RepositoryAnalytics) *ServiceAnalytics {
	return &ServiceAnalytics{
		repo: repo,
	}
}
