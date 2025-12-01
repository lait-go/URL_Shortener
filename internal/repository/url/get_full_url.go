package url

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (r *RepositoryURL) GetFullURL(ctx context.Context, short_url string) (string, error) {
	query := `
		SELECT full_url
		FROM url
		WHERE short_url = $1;
	`
	var fullURL string
	err := r.db.QueryRowContext(ctx, query, short_url).Scan(&fullURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrFullURLNotFound
		}

		return "", fmt.Errorf("repository/get_full_url.go - failed to get full url - %w", err)
	}

	return fullURL, nil
}
