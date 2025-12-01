package url

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/avraam311/url-shortener/internal/models/domain"
	"github.com/avraam311/url-shortener/internal/models/dto"
)

const urlLength = 13

func (s *ServiceURL) CreateShortURL(ctx context.Context, fullURL *dto.FullURL) (string, error) {
	var shortURL string
	var err error
	// to avoid collisions
	uniqueShortURL := false
	for !uniqueShortURL {
		shortURL, err = generateShortURL(fullURL.URL)
		if err != nil {
			return "", err
		}
		unique, err := s.repo.CheckIfShortURLIsUnique(ctx, shortURL)
		if err != nil {
			return "", fmt.Errorf("service/create_short_url.go - %w", err)
		}
		uniqueShortURL = unique
	}

	url := domain.URL{
		FullURL:  fullURL.URL,
		ShortURL: shortURL,
	}

	shortURL, err = s.repo.SaveShortURL(ctx, &url)
	if err != nil {
		return "", fmt.Errorf("service/create_short_url.go - %w", err)
	}

	return shortURL, nil
}

func generateShortURL(fullURL string) (string, error) {
	salt := make([]byte, 8)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	combined := append([]byte(fullURL), salt...)
	hash := sha256.Sum256(combined)
	encoded := base64.URLEncoding.EncodeToString(hash[:])
	shortURL := encoded[:urlLength]

	return shortURL, nil
}
