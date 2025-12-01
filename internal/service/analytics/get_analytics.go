package analytics

import (
	"context"
	"fmt"

	"github.com/avraam311/url-shortener/internal/models/db"
)

func (s *ServiceAnalytics) GetAnalytics(ctx context.Context, shortURL string) ([]*db.Analytics, error) {
	analytics, err := s.repo.GetAnalytics(ctx, shortURL)
	if err != nil {
		return nil, fmt.Errorf("service/get_analytics - %w", err)
	}

	return analytics, nil
}
