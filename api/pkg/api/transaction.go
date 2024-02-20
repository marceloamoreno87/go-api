package api

import (
	"database/sql"
)

type Transaction struct {
	dbConn *sql.DB
}

func NewTransaction(dbConn *sql.DB) *Transaction {
	return &Transaction{
		dbConn: dbConn,
	}
}

func (t *Transaction) Begin(fn func(*sql.Tx)) (tx *sql.Tx, err error) {
	tx, err = t.dbConn.Begin()
	if err != nil {
		return
	}
	fn(tx)
	return
}

func (t *Transaction) Commit(tx *sql.Tx) (err error) {
	err = tx.Commit()
	return
}

func (t *Transaction) Rollback(tx *sql.Tx) (err error) {
	err = tx.Rollback()
	return
}
