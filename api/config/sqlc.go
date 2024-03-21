package config

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type SQLCInterface interface {
	GetDbQueries() *db.Queries
	GetDbConn() *sql.DB
	GetTx() *sql.Tx
	Begin()
	Rollback()
	Commit()
	SetCtx(ctx context.Context)
	GetCtx() context.Context
}

type SQLC struct {
	dbConn    *sql.DB
	dbQueries *db.Queries
	tx        *sql.Tx
	ctx       context.Context
}

func NewSqlc(DB DatabaseInterface) SQLCInterface {
	return &SQLC{
		dbConn:    DB.GetDbConn(),
		dbQueries: db.New(DB.GetDbConn()),
		tx:        nil,
		ctx:       nil,
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

func (t *SQLC) Begin() {
	t.tx, _ = t.dbConn.Begin()
}

func (t *SQLC) Rollback() {
	t.tx.Rollback()
}

func (t *SQLC) Commit() {
	t.tx.Commit()
}

func (t *SQLC) SetCtx(ctx context.Context) {
	t.ctx = ctx
}

func (t *SQLC) GetCtx() context.Context {
	return t.ctx
}
