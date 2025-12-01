package url

import (
	"context"

	"github.com/avraam311/url-shortener/internal/models/dto"

	"github.com/go-playground/validator/v10"
)

//go:generate mockgen -source=handler.go -destination=../../../mocks/url/handler.go
type ServiceURL interface {
	CreateShortURL(context.Context, *dto.FullURL) (string, error)
	GetFullURL(context.Context, string) (string, error)
}

type HandlerURL struct {
	service   ServiceURL
	validator *validator.Validate
}

func NewHandler(service ServiceURL, validator *validator.Validate) *HandlerURL {
	return &HandlerURL{
		service:   service,
		validator: validator,
	}
}
