package helper

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetLimitAndOffset(r *http.Request) (limit int32, offset int32) {
	limit = StrToInt32(chi.URLParam(r, "limit"))
	offset = StrToInt32(chi.URLParam(r, "offset"))

	if limit == 0 {
		limit = 10
	}

	if offset == 0 {
		offset = 0
	}

	return
}

func GetID(r *http.Request) (id int32) {
	id = StrToInt32(chi.URLParam(r, "id"))
	return
}
