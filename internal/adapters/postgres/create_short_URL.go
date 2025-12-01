package postgres

import (
	"context"
	"crudl/internal/domain"
	"fmt"
)

func (p *Pool) Set(ctx context.Context, data *domain.ShortURL) error {
	tx, err := p.DB.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin tx: %w", err)
	}
	defer tx.Rollback()

    if _, err := tx.NamedExecContext(ctx, `
        INSERT INTO shortened_urls (short_key, original_url, created_at) 
        VALUES (:short_code, :original_url, :created_at)`, data, ); err != nil {
		return fmt.Errorf("error inserting short URL: %w", err)
	}

	return tx.Commit()
}

func (p *Pool) SetToAnalitics(ctx context.Context, data *domain.ClickAnalytics) error {
    tx, err := p.DB.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin tx: %w", err)
	}
	defer tx.Rollback()

    id, err := p.getIdShortURL(ctx, data.ShortCode)
    if err != nil{
        return err
    }

    if _, err := tx.NamedExecContext(ctx, `
        INSERT INTO url_clicks (short_url_id, user_agent, ip_address, clicked_at) 
        VALUES ($1, :user_agent, :ip, :timestamp)`, id); err != nil {
		return fmt.Errorf("error inserting short URL: %w", err)
	}

	return tx.Commit()
}
