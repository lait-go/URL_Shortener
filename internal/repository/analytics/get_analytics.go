package analytics

import (
	"context"
	"fmt"

	"github.com/avraam311/url-shortener/internal/models/db"
)

func (r *RepositoryAnalytics) GetAnalytics(ctx context.Context, shortURL string) ([]*db.Analytics, error) {
	query := `
		SELECT id, short_url, ip, user_agent, time
		FROM analytics
		WHERE short_url = $1;
	`

	rows, err := r.db.QueryContext(ctx, query, shortURL)
	if err != nil {
		return nil, fmt.Errorf("repository/get_analytics.go - failed to get analytics - %w", err)
	}
	defer rows.Close()

	analytics := make([]*db.Analytics, 0)
	for rows.Next() {
		a := &db.Analytics{}
		if err := rows.Scan(&a.ID, &a.ShortURL, &a.Ip, &a.UserAgent, &a.Time); err != nil {
			return nil, fmt.Errorf("repository/get_analytics.go - failed to scan analytics rows - %w", err)
		}

		analytics = append(analytics, a)
	}

	return analytics, nil
}
