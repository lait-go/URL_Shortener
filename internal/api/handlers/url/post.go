package url

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/avraam311/url-shortener/internal/api/handlers"
	"github.com/avraam311/url-shortener/internal/models/dto"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *HandlerURL) CreateShortURL(c *ginext.Context) {
	var FullURL *dto.FullURL

	if err := json.NewDecoder(c.Request.Body).Decode(&FullURL); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to decode request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error()))
		return
	}

	if err := h.validator.Struct(FullURL); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to validate request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("validation error: %s", err.Error()))
		return
	}

	shortURL, err := h.service.CreateShortURL(c.Request.Context(), FullURL)
	if err != nil {
		zlog.Logger.Error().Err(err).Interface("message", FullURL.URL).Msg("failed to create short url")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.Created(c.Writer, shortURL)
}
