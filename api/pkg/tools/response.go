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
	Data       interface{} `json:"data"`
	StatusCode int         `json:"status_code"`
}

type ResponseError struct {
	Msg        string `json:"msg"`
	StatusCode int    `json:"status_code"`
	CodeError  string `json:"code_error"`
}

func (rt *ResponseTools) NewResponse(
	data interface{},
	statusCode int,
) Response {
	return Response{
		Data:       data,
		StatusCode: statusCode,
	}
}

func (rt *ResponseTools) NewResponseError(
	msg string,
	statusCode int,
	codeError string,
) ResponseError {
	return ResponseError{
		Msg:        msg,
		StatusCode: statusCode,
		CodeError:  codeError,
	}
}

func (rt *ResponseTools) ResponseJSON(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func (rt *ResponseTools) ResponseErrorJSON(w http.ResponseWriter, responseError ResponseError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseError.StatusCode)
	slog.Error(responseError.Msg, "code_error", responseError.CodeError)
	json.NewEncoder(w).Encode(responseError)
}

func (rt *ResponseTools) ToJson(data interface{}) string {
	json, err := json.Marshal(data)
	if err != nil {
		slog.Error("err", err)
	}
	return string(json)
}
