package url

import (
	"context"
	"fmt"

	"github.com/avraam311/url-shortener/internal/models/domain"
)

func (r *RepositoryURL) SaveShortURL(ctx context.Context, url *domain.URL) (string, error) {
	query := `
		WITH ins AS (
			INSERT INTO url (full_url, short_url)
			VALUES ($1, $2)
			ON CONFLICT (full_url) DO NOTHING
			RETURNING short_url
		)
		SELECT short_url FROM ins
		UNION ALL
		SELECT short_url FROM url WHERE full_url = $1
		LIMIT 1;
	`
	var shortURL string
	err := r.db.QueryRowContext(ctx, query, url.FullURL, url.ShortURL).Scan(&shortURL)
	if err != nil {
		return "", fmt.Errorf("repository/save_short_url.go - failed to save short url - %w", err)
	}

	return shortURL, nil
}
