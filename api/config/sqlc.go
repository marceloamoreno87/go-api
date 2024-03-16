package config

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type SQLCInterface interface {
	GetDbQueries() *db.Queries
	GetDbConn() *sql.DB
}

type SQLC struct {
	dbConn    *sql.DB
	dbQueries *db.Queries
}

func NewSqlc(DB DatabaseInterface) SQLCInterface {
	return &SQLC{
		dbConn:    DB.GetDbConn(),
		dbQueries: db.New(DB.GetDbConn()),
	}
}

func (t *SQLC) GetDbConn() *sql.DB {
	return t.dbConn
}

func (t *SQLC) GetDbQueries() *db.Queries {
	return t.dbQueries
}
