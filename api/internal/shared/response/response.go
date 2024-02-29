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
	Data       interface{} `json:"data"`
	StatusCode int         `json:"status_code"`
}

type ResponseError struct {
	Msg       string `json:"msg"`
	CodeError string `json:"code_error"`
}

func (rt *Responses) NewResponse(
	data interface{},
	statusCode int,
) Response {
	return Response{
		Data:       data,
		StatusCode: statusCode,
	}
}

func (rt *Responses) NewResponseError(
	msg string,
	statusCode int,
	codeError string,
) ResponseError {
	return ResponseError{
		Msg:       msg,
		CodeError: codeError,
	}
}

func (rt *Responses) SendResponse(
	w http.ResponseWriter,
	response Response,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
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
