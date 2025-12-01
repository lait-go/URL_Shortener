package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/avraam311/url-shortener/internal/models/dto"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

type ServiceAnalytics interface {
	SaveAnalytics(context.Context, *dto.Analytics) error
}

func CORSMiddleware() ginext.HandlerFunc {
	return func(c *ginext.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func AnalyticsMiddleware(repoAnalytics ServiceAnalytics) ginext.HandlerFunc {
	return func(c *ginext.Context) {
		shortURL := c.Param("short_url")
		ip := c.ClientIP()
		userAgent := c.GetHeader("User-Agent")
		time := time.Now()

		c.Next()

		analytics := dto.Analytics{
			ShortURL:  shortURL,
			Ip:        ip,
			UserAgent: userAgent,
			Time:      time,
		}
		if err := repoAnalytics.SaveAnalytics(c.Request.Context(), &analytics); err != nil {
			zlog.Logger.Warn().Err(err).Msg("failed to save analytics")
			return
		}
	}
}
