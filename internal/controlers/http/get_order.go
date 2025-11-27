package http

import (
	"crudl/internal/domain"
	"crudl/pkg/logger"
	"crudl/pkg/render"
	"encoding/json"
	"net/http"
)

// GetSub godoc
// @Summary Получить подписку
// @Description Возвращает данные подписки по ID
// @Tags subscriptions
// @Accept  json
// @Produce  json
// @Param   id  path  string  true  "ID подписки"
// @Success 200 {object} domain.Order
// @Failure 400 {object} map[string]string "invalid ID"
// @Failure 404 {object} map[string]string "not found"
// @Failure 500 {object} map[string]string "internal server error"
// @Router /sub/{id} [get]
func (h *Handlers) PostShortner(w http.ResponseWriter, r *http.Request) {
	var shortURLRequest domain.CreateShortURLRequest

	if err := json.NewDecoder(r.Body).Decode(&shortURLRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Default().InfoContext(r.Context(), "start working: ", shortURLRequest)

	shortURL, err := h.profileService.SaveURL(r.Context(), shortURLRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Default().InfoContext(r.Context(), "ending working: ", shortURL)

	render.JSON(w, *shortURL, http.StatusOK)
}
