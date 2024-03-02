package response

import (
	"encoding/json"
	"net/http"
)

type ResponsesInterface interface {
	NewResponse(data interface{}, statusCode int) Response
	NewResponseError(msg string, statusCode int, codeError string) ResponseError
	SendResponse(w http.ResponseWriter, response Response)
	SendResponseError(w http.ResponseWriter, responseError ResponseError)
}

type Responses struct {
	Response
	ResponseError
}

type Response struct {
	Data any `json:"data"`
}

type ResponseError struct {
	Err string `json:"err"`
}

func (rt *Responses) NewResponse(
	data any,
) Response {
	return Response{
		Data: data,
	}
}

func (rt *Responses) NewResponseError(
	err string,
) ResponseError {
	return ResponseError{
		Err: err,
	}
}

func (rt *Responses) SendResponse(
	w http.ResponseWriter,
	response Response,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (rt *Responses) SendResponseError(
	w http.ResponseWriter,
	responseError ResponseError,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(responseError)
}
