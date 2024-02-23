package tools

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type HandlerToolsInterface interface {
	GetLimitOffsetFromURL(r *http.Request) (int32, int32, error)
	GetIDFromURL(r *http.Request) (int32, error)
	ResponseToolsInterface
}

type HandlerTools struct {
	*ResponseTools
}

func NewHandlerTools() *HandlerTools {
	return &HandlerTools{
		ResponseTools: &ResponseTools{},
	}
}

func (ht *HandlerTools) GetLimitOffsetFromURL(r *http.Request) (limitInt int32, offsetInt int32, err error) {

	limitInt = int32(10)
	offsetInt = int32(0)

	limit := r.URL.Query().Get("limit")
	if limit != "" {
		limitIntTemp, _ := strconv.Atoi(limit)
		limitInt = int32(limitIntTemp)
	}

	if limitInt < 0 {
		return 0, 0, errors.New("limit must be greater than 0")
	}

	offset := r.URL.Query().Get("offset")
	if offset != "" {
		offsetIntTemp, _ := strconv.Atoi(offset)
		offsetInt = int32(offsetIntTemp)
	}

	if offsetInt < 0 {
		return 0, 0, errors.New("offset must be greater than 0")
	}

	limitInt = int32(limitInt)
	offsetInt = int32(offsetInt)
	return
}

func (ht *HandlerTools) GetIDFromURL(r *http.Request) (idInt int32, err error) {
	id := chi.URLParam(r, "id")

	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		panic(err)
	}
	idInt = int32(i)
	return
}
