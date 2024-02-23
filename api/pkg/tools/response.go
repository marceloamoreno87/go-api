package tools

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ResponseToolsInterface interface {
	ResponseJSON(w http.ResponseWriter, response Response)
	ResponseErrorJSON(w http.ResponseWriter, responseError ResponseError)
	ToJson(data interface{}) string
	NewResponse(data interface{}, statusCode int) Response
	NewResponseError(msg string, statusCode int, codeError string) ResponseError
}

type ResponseTools struct {
	Response
	ResponseError
}

type Response struct {
	data       interface{}
	statusCode int
}

type ResponseError struct {
	msg        string
	statusCode int
	codeError  string
}

func (rt *ResponseTools) NewResponse(
	data interface{},
	statusCode int,
) Response {
	return Response{
		data:       data,
		statusCode: statusCode,
	}
}

func (rt *ResponseTools) NewResponseError(
	msg string,
	statusCode int,
	codeError string,
) ResponseError {
	return ResponseError{
		msg:        msg,
		statusCode: statusCode,
		codeError:  codeError,
	}
}

func (rt *ResponseTools) ResponseJSON(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.statusCode)
	json.NewEncoder(w).Encode(rt.ToJson(response))
}

func (rt *ResponseTools) ResponseErrorJSON(w http.ResponseWriter, responseError ResponseError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseError.statusCode)
	slog.Error(responseError.msg, "code_error", responseError.codeError)
	json.NewEncoder(w).Encode(rt.ToJson(responseError))
}

func (rt *ResponseTools) ToJson(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(b)
}
