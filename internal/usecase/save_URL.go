package usecase

import (
	"context"
	"crudl/internal/domain"
	"math/rand/v2"
	"strconv"
	"time"
)

func (p *Profile) SaveURL(ctx context.Context, data *domain.ShortURLRequest) (*domain.ShortURLRequest, error) {
	shortData := &domain.ShortURL{
		ShortCode:   data.ShortCode,
		OriginalURL: data.OriginalURL,
		CreatedAt:   time.Now(),
	}

	if data.ShortCode == "" {
		data.ShortCode = p.checkURLToCache(shortData)
	}

	if err := p.Cache.Set(ctx, shortData); err != nil {
		return nil, err
	}

	err := p.Postgres.Set(ctx, shortData)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// TODO: реализовать поиск в бд
func (p *Profile) checkURLToCache(data *domain.ShortURL) string {
	var shortCode string
	for {
		shortCode = createShortURL(data.OriginalURL)

		if err := p.Cache.Set(context.Background(), data); err != nil {
			if err.Error() == domain.NOTFOUND {
				continue
			}
		}
		break
	}

	return shortCode
}

func createShortURL(URL string) string {
	return URL[6:rand.IntN(8)] + strconv.Itoa(rand.IntN(256))
}
