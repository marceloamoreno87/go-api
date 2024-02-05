package api

import "net/http"

type HandlerToolsInterface interface {
	GetLimitOffsetFromURL(r *http.Request) (int32, int32, error)
	GetIDFromURL(r *http.Request) (int32, error)
	ResponseJSON(w http.ResponseWriter, data interface{})
	ResponseErrorJSON(w http.ResponseWriter, responseError ResponseError)
}
