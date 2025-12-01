package http

import (
	"crudl/internal/domain"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

func (h *Handlers) TransitionToURL(w http.ResponseWriter, r *http.Request) {
	Analytic := domain.ClickAnalytics{
		ShortCode: chi.URLParam(r, "short_url"),
		Timestamp: time.Now(),
		UserAgent: r.UserAgent(),
		IP:        r.RemoteAddr,
	}

	res, err := h.profileService.GetURLFromShortURL(r.Context(), Analytic)
	if err != nil {
		handleUseCaseError(w, r, err)
		return
	}

	http.Redirect(w, r, res.OriginalURL, http.StatusOK)
}
