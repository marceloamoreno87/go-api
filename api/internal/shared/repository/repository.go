package repository

import (
	"database/sql"

	"github.com/marceloamoreno/goapi/config"
	"github.com/marceloamoreno/goapi/internal/shared/db"
)

type RepositoryInterface interface {
	GetDbQueries() *db.Queries
	GetDbConn() *sql.DB
	Begin() (err error)
	Commit() (err error)
	Rollback() (err error)
	GetTx() *sql.Tx
}

type Repository struct {
	dbConn    *sql.DB
	dbQueries *db.Queries
	tx        *sql.Tx
}

func NewRepository(DB config.DatabaseInterface) *Repository {
	return &Repository{
		dbConn:    DB.GetDbConn(),
		dbQueries: db.New(DB.GetDbConn()),
		tx:        nil,
	}
}

func (t *Repository) GetDbConn() *sql.DB {
	return t.dbConn
}

func (t *Repository) GetDbQueries() *db.Queries {
	return t.dbQueries
}

func (t *Repository) Begin() (err error) {
	tx, err := t.dbConn.Begin()
	if err != nil {
		return
	}
	t.tx = tx
	return
}

func (t *Repository) Commit() (err error) {
	err = t.tx.Commit()
	return
}

func (t *Repository) Rollback() (err error) {
	err = t.tx.Rollback()
	return
}

func (t *Repository) GetTx() *sql.Tx {
	return t.tx
}
