package api

import (
	"net/http"
)

type Response struct {
	Data       interface{} `json:"data"`
	StatusCode int         `json:"status_code"`
}

func NewResponse(data interface{}, statusCode int) Response {
	return Response{
		Data:       data,
		StatusCode: statusCode,
	}
}

type ResponseError struct {
	Msg        string `json:"msg"`
	StatusCode int    `json:"status_code"`
	CodeError  string `json:"code_error"`
}

func NewResponseError(msg string, statusCode int, codeError string) ResponseError {
	return ResponseError{
		Msg:        msg,
		StatusCode: statusCode,
		CodeError:  codeError,
	}
}

func NewResponseErrorDefault(msg string) ResponseError {
	return ResponseError{
		Msg:        msg,
		StatusCode: http.StatusBadRequest,
		CodeError:  "BAD_REQUEST",
	}
}

var (
	NOT_AUTHORIZED = NewResponseError("Not Authorized", http.StatusUnauthorized, "NOT_AUTHORIZED")
	NOT_FOUND      = NewResponseError("Not Found", http.StatusNotFound, "NOT_FOUND")
	BAD_REQUEST    = NewResponseError("Bad Request", http.StatusBadRequest, "BAD_REQUEST")
)
