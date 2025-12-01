package usecase

import (
	"context"
	"crudl/internal/domain"
	"crudl/pkg/logger"
)

func (p *Profile) GetURLFromShortURL(ctx context.Context, analitic domain.ClickAnalytics) (*domain.ShortURL, error) {
	if analitic.ShortCode == "" {
		return nil, &domain.NotFoundError{}
	}

	dataCashe, err := p.Cache.Get(ctx, analitic.ShortCode)
	if err != nil {
		return nil, err
	}
	
	res := &domain.ShortURL{
		OriginalURL: dataCashe.OriginalURL,
		ShortCode:   dataCashe.ShortCode,
	}

	if res == nil {
		res, err = p.Postgres.Get(ctx, analitic.ShortCode)
		if err != nil {
			return nil, err
		}
	}

	go p.AddToAnalitic(ctx, res)

	return res, nil
}

func (p *Profile) AddToAnalitic(ctx context.Context, data *domain.ShortURL) {
	if err := p.Cache.Set(ctx, data); err != nil {
		logger.Default().ErrorContext(ctx, err.Error(), err)
	}

	if err := p.Postgres.Set(ctx, data); err != nil {
		logger.Default().ErrorContext(ctx, err.Error(), err)
	}
}
