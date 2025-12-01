package postgres

import (
	"context"
	"crudl/internal/domain"
)

func (p *Pool) Get(ctx context.Context, shortURL string) (*domain.ShortURL, error) {
	rows, err := p.DB.QueryContext(ctx, `SELECT * FROM shortened_urls WHERE short_key = $1`, shortURL)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	sub := &domain.ShortURL{}

	for rows.Next() {
		if err := rows.Scan(
			&sub.ShortCode,
			&sub.OriginalURL,
			&sub.CreatedAt,
		); err != nil {
			return nil, err
		}
	}

	if sub.OriginalURL == "" {
		return nil, &domain.NotFoundError{
			Message: domain.NOTFOUND,
		}
	}

	return sub, nil
}

func (p *Pool) getIdShortURL(ctx context.Context, shortURL string) (int, error) {
	var id int
	if err := p.DB.QueryRowContext(ctx, `SELECT id FROM shortened_urls WHERE short_key = $1`, shortURL).Scan(&id); err != nil {
		return 0, &domain.NotFoundError{
			Message: domain.NOTFOUND,
		}
	}

	return id, nil
}
