package usecase

import (
	"context"
	"crudl/internal/domain"
	"math/rand/v2"
	"strconv"
	"time"
)

func (p *Profile) SaveURL(ctx context.Context, data domain.CreateShortURLRequest) (*domain.ShortURL, error) {
	if data.CustomAlias == "" {
		data.CustomAlias = p.checkURLToCache(data)
	}

	if err := p.Cache.Set(ctx, data); err != nil {
		return &domain.ShortURL{}, err
	}

	short := &domain.ShortURL{
		OriginalURL: data.OriginalURL,
		ShortCode:   data.CustomAlias,
		CreatedAt:   time.Now(),
	}

	err := p.Postgres.Set(ctx, *short)
	if err != nil {
		return &domain.ShortURL{}, err
	}

	return short, nil
}

// TODO: реализовать поиск в бд
func (p *Profile) checkURLToCache(data domain.CreateShortURLRequest) string {
	var shortCode string
	for {
		shortCode = createShortURL(data.OriginalURL)

		if err := p.Cache.Set(context.Background(), data); err != nil {
			if err.Error() == domain.ErrorExists {
				continue
			}
		}
		break
	}

	return shortCode
}

func createShortURL(URL string) string {
	return URL[:rand.IntN(8)] + strconv.Itoa(rand.IntN(256))
}
