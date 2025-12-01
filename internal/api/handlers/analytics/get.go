package analytics

import (
	"fmt"
	"net/http"

	"github.com/avraam311/url-shortener/internal/api/handlers"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *HandlerAnalytics) GetAnalytics(c *ginext.Context) {
	shortURL := c.Param("short_url")

	if shortURL == "" {
		zlog.Logger.Error().Msg("empty short url")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("empty short url"))
		return
	}

	analytics, err := h.service.GetAnalytics(c.Request.Context(), shortURL)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to get analytics")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.OK(c.Writer, analytics)
}
