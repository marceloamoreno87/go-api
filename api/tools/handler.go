package tools

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const Error string = "Error"
const Success string = "Success"
const Info string = "Info"

type HandlerToolsInterface interface {
	GetLimitOffsetFromURL(r *http.Request) (int32, int32, error)
	GetIDFromURL(r *http.Request) (int64, error)
	ResponseJSON(w http.ResponseWriter, statusCode int, data interface{})
	ResponseErrorJSON(w http.ResponseWriter, statusCode int, err interface{})
}

type Response struct {
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Err interface{} `json:"err"`
}

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

func (h *HandlerTools) GetIDFromURL(r *http.Request) (idInt int64, err error) {

	id := chi.URLParam(r, "id")
	if id == "" {
		return 0, errors.New("id is required")
	}

	idInt, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, errors.New("id must be a number")
	}

	return
}

func (h *HandlerTools) ResponseJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&Response{
		Data: data,
	})
}

func (h *HandlerTools) ResponseErrorJSON(w http.ResponseWriter, statusCode int, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&ResponseError{
		Err: err,
	})
}
