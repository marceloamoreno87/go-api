package config

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type SQLCInterface interface {
	GetDbQueries() *db.Queries
	GetDbConn() *sql.DB
	GetTx() *sql.Tx
	SetTx(tx *sql.Tx)
}

type SQLC struct {
	dbConn    *sql.DB
	dbQueries *db.Queries
	tx        *sql.Tx
}

func NewSqlc(DB DatabaseInterface) SQLCInterface {
	return &SQLC{
		dbConn:    DB.GetDbConn(),
		dbQueries: db.New(DB.GetDbConn()),
		tx:        nil,
	}
}

func (t *SQLC) GetDbConn() *sql.DB {
	return t.dbConn
}

func (t *SQLC) GetDbQueries() *db.Queries {
	return t.dbQueries
}

func (t *SQLC) GetTx() *sql.Tx {
	return t.tx
}

func (t *SQLC) SetTx(tx *sql.Tx) {
	t.tx = tx
}
