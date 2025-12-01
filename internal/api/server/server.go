package server

import (
	"context"
	"net/http"

	"github.com/wb-go/wbf/ginext"

	"github.com/avraam311/url-shortener/internal/api/handlers/analytics"
	"github.com/avraam311/url-shortener/internal/api/handlers/url"
	"github.com/avraam311/url-shortener/internal/middlewares"
	"github.com/avraam311/url-shortener/internal/models/dto"
)

type ServiceAnalytics interface {
	SaveAnalytics(context.Context, *dto.Analytics) error
}

func NewRouter(handlerURL *url.HandlerURL, handlerAnalytics *analytics.HandlerAnalytics, serviceAnalytics ServiceAnalytics, ginMode string) *ginext.Engine {
	e := ginext.New(ginMode)

	e.Use(middlewares.CORSMiddleware())
	e.Use(ginext.Logger())
	e.Use(ginext.Recovery())

	api := e.Group("/api/url-shortener")
	{
		api.POST("/shorten", handlerURL.CreateShortURL)
		api.GET("/:short_url", middlewares.AnalyticsMiddleware(serviceAnalytics), handlerURL.GoToShortUrl)
		api.GET("/analytics/:short_url", handlerAnalytics.GetAnalytics)
	}

	return e
}

func NewServer(addr string, router *ginext.Engine) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
