package url

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (r *RepositoryURL) CheckIfShortURLIsUnique(ctx context.Context, shortURL string) (bool, error) {
	query := `
		SELECT id
		FROM url
		WHERE short_url = $1;
	`
	var shortURLID int
	err := r.db.QueryRowContext(ctx, query, shortURL).Scan(&shortURLID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}

		return false, fmt.Errorf("repository/check_if_short_url_is_unique - failed to check - %w", err)
	}

	return false, nil
}
