package analytics

import (
	"context"
	"fmt"

	"github.com/avraam311/url-shortener/internal/models/dto"
)

func (r *RepositoryAnalytics) SaveAnalytics(ctx context.Context, analytics *dto.Analytics) error {
	query := `
		INSERT INTO analytics (
			short_url, ip, user_agent, time
		) VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.ExecContext(ctx, query, analytics.ShortURL, analytics.Ip, analytics.UserAgent, analytics.Time)
	if err != nil {
		return fmt.Errorf("repository/save_analytics.go - failed to save analytics - %w", err)
	}

	return nil
}
