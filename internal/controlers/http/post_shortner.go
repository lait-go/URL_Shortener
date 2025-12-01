package http

import (
	"crudl/internal/domain"
	"crudl/pkg/logger"
	"crudl/pkg/render"
	"encoding/json"
	"net/http"
)

func (h *Handlers) PostShortner(w http.ResponseWriter, r *http.Request) {
	var shortURLRequest *domain.ShortURLRequest

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
