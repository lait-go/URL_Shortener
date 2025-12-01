package url

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/avraam311/url-shortener/internal/api/handlers"
	"github.com/avraam311/url-shortener/internal/repository/url"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *HandlerURL) GoToShortUrl(c *ginext.Context) {
	shortURL := c.Param("short_url")

	if shortURL == "" {
		zlog.Logger.Error().Msg("empty short url")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("empty short url"))
		return
	}

	fullURL, err := h.service.GetFullURL(c.Request.Context(), shortURL)
	if err != nil {
		if errors.Is(err, url.ErrFullURLNotFound) {
			zlog.Logger.Error().Err(err).Str("short_url", shortURL).Msg("full url not found")
			handlers.Fail(c.Writer, http.StatusNotFound, url.ErrFullURLNotFound)
			return
		}

		zlog.Logger.Error().Err(err).Msg("failed to get full url")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	c.Redirect(http.StatusFound, fullURL)
}
