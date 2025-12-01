package http

import (
	"crudl/internal/domain"
	"crudl/pkg/logger"
	"crudl/pkg/render"
	"net/http"
)

func newErrorResponse(code string, message string) domain.ErrorResponse {
	var errResp domain.ErrorResponse
	errResp.Error.Code = code
	errResp.Error.Message = message
	return errResp
}

func handleUseCaseError(w http.ResponseWriter, r *http.Request, err error) {
	var errResp domain.ErrorResponse
	var status int

	logger.Default().ErrorContext(r.Context(),
		"usecase error", err,
		"path", r.URL.Path,
		"method", r.Method,
	)

	switch e := err.(type) {
	case *domain.KeyExistsError:
		errResp.Error.Code = domain.KEYEXISTS
		errResp.Error.Message = e.Error()
		status = http.StatusBadRequest

	case *domain.NotFoundError:
		errResp.Error.Code = domain.NOTFOUND
		errResp.Error.Message = e.Error()
		status = http.StatusBadRequest

	default: //TODO: придумать что использовать вместо NOTFOUND
		errResp.Error.Code = domain.NOTFOUND
		errResp.Error.Message = "internal server error"
		status = http.StatusInternalServerError
	}

	renderError(w, errResp, status)
}

func renderError(w http.ResponseWriter, errResp domain.ErrorResponse, status int) {
	render.JSON(w, errResp, status)
}
