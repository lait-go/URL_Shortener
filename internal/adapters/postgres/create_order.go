package postgres

import (
	"context"
	"crudl/internal/domain"
	"fmt"
)

func (p *Pool) Set(ctx context.Context, data domain.ShortURL) error {
    tx, err := p.DB.BeginTxx(ctx, nil)
    if err != nil {
        return fmt.Errorf("failed to begin tx: %w", err)
    }
    defer tx.Rollback()
	fmt.Println(data)
    if _, err := tx.NamedExecContext(ctx, `
        INSERT INTO shortened_urls (short_key, original_url, created_at) 
        VALUES (:short_key, :original_url, :created_at)`, data); err != nil {
        return fmt.Errorf("error inserting short URL: %w", err)
    }

    return tx.Commit()
}
