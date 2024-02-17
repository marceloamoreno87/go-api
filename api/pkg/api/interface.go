package api

import (
	"net/http"
)

type HandlerToolsInterface interface {
	GetLimitOffsetFromURL(r *http.Request) (int32, int32, error)
	GetIDFromURL(r *http.Request) (int32, error)
	ResponseJSON(w http.ResponseWriter, data interface{})
	ResponseErrorJSON(w http.ResponseWriter, responseError ResponseError)
}

type DatabaseTransaction interface {
	BeginTx() (err error)
	CommitTx() (err error)
	RollbackTx() (err error)
}

// type DatabaseTransaction interface {
// 	BeginTx(ctx context.Context, options *sql.TxOptions) (tx *sql.Tx, err error)
// 	CommitTx(tx *sql.Tx) (err error)
// 	RollbackTx(tx *sql.Tx) (err error)
// }
