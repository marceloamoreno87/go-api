package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type HandlerTools struct {
	dbConn *sql.DB
}

func NewHandlerTools(dbConn *sql.DB) *HandlerTools {
	return &HandlerTools{
		dbConn: dbConn,
	}
}

func (h *HandlerTools) GetLimitOffsetFromURL(r *http.Request) (int32, int32, error) {

	limitInt := 10
	offsetInt := 0

	limit := r.URL.Query().Get("limit")
	if limit != "" {
		limitInt, _ = strconv.Atoi(limit)
	}

	if limitInt < 0 {
		return 0, 0, errors.New("limit must be greater than 0")
	}

	offset := r.URL.Query().Get("offset")
	if offset != "" {
		offsetInt, _ = strconv.Atoi(offset)
	}

	if offsetInt < 0 {
		return 0, 0, errors.New("offset must be greater than 0")
	}

	return int32(limitInt), int32(offsetInt), nil
}

func (h *HandlerTools) GetIDFromURL(r *http.Request) (idInt int32, err error) {
	id := chi.URLParam(r, "id")

	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		panic(err)
	}
	return int32(i), err
}

func (h *HandlerTools) ResponseJSON(w http.ResponseWriter, data interface{}) {
	response := NewResponse(data, http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func (h *HandlerTools) ResponseErrorJSON(w http.ResponseWriter, responseError ResponseError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseError.StatusCode)
	slog.Error(responseError.Msg, "code_error", responseError.CodeError)
	json.NewEncoder(w).Encode(responseError)
}
