package tools

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type HandlerTools struct {
}

func NewHandlerTools() *HandlerTools {
	return &HandlerTools{}

}

func (h *HandlerTools) GetLimitOffsetFromURL(r *http.Request) (int32, int32, error) {

	limitInt := 10
	offsetInt := 0

	limit := r.URL.Query().Get("limit")
	if limit != "" {
		limitInt, _ = strconv.Atoi(limit)
		if limitInt < 0 {
			return 0, 0, errors.New("limit must be greater than 0")
		}
	}

	offset := r.URL.Query().Get("offset")
	if offset != "" {
		offsetInt, _ = strconv.Atoi(offset)
		if offsetInt < 0 {
			return 0, 0, errors.New("offset must be greater than 0")
		}
	}

	return int32(limitInt), int32(offsetInt), nil
}

func (h *HandlerTools) GetIDFromURL(r *http.Request) (int64, error) {

	id := chi.URLParam(r, "id")
	if id == "" {
		return 0, errors.New("id is required")
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("id must be a number")
	}

	return int64(idInt), nil
}