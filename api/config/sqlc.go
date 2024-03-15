package config

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type SQLCInterface interface {
	GetDbQueries() *db.Queries
	GetDbConn() *sql.DB
	Begin() (err error)
	Commit() (err error)
	Rollback() (err error)
	GetTx() *sql.Tx
	Close()
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

func (t *SQLC) Begin() (err error) {
	tx, err := t.dbConn.Begin()
	if err != nil {
		return
	}
	t.tx = tx
	return
}

func (t *SQLC) Commit() (err error) {
	err = t.tx.Commit()
	return
}

func (t *SQLC) Rollback() (err error) {
	err = t.tx.Rollback()
	return
}

func (t *SQLC) GetTx() *sql.Tx {
	return t.tx
}

func (t *SQLC) Close() {
	t.dbConn.Close()
}
