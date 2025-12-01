package analytics

import (
	"context"

	"github.com/avraam311/url-shortener/internal/models/db"

	"github.com/go-playground/validator/v10"
)

//go:generate mockgen -source=handler.go -destination=../../../mocks/analytics/handler.go
type ServiceAnalytics interface {
	GetAnalytics(context.Context, string) ([]*db.Analytics, error)
}

type HandlerAnalytics struct {
	service   ServiceAnalytics
	validator *validator.Validate
}

func NewHandler(service ServiceAnalytics, validator *validator.Validate) *HandlerAnalytics {
	return &HandlerAnalytics{
		service:   service,
		validator: validator,
	}
}
