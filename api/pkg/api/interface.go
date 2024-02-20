package api

import (
	"database/sql"
	"net/http"
)

type HandlerToolsInterface interface {
	GetLimitOffsetFromURL(r *http.Request) (int32, int32, error)
	GetIDFromURL(r *http.Request) (int32, error)
	ResponseJSON(w http.ResponseWriter, data interface{})
	ResponseErrorJSON(w http.ResponseWriter, responseError ResponseError)
}

type TransactionInterface interface {
	Begin(fn func(*sql.Tx)) (tx *sql.Tx, err error)
	Commit(tx *sql.Tx) (err error)
	Rollback(tx *sql.Tx) (err error)
}
