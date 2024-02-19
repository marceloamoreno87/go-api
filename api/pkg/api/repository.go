package api

import (
	"context"
	"database/sql"
)

type DefaultRepository struct {
	DbConn *sql.DB
}

func NewDefaultRepository(dbConn *sql.DB) *DefaultRepository {
	return &DefaultRepository{
		DbConn: dbConn,
	}
}

func (dr *DefaultRepository) BeginTx(ctx context.Context, options *sql.TxOptions) (tx *sql.Tx, err error) {
	tx, err = dr.DbConn.BeginTx(ctx, options)
	return
}

func (dr *DefaultRepository) CommitTx(tx *sql.Tx) (err error) {
	err = tx.Commit()
	return
}

func (dr *DefaultRepository) RollbackTx(tx *sql.Tx) (err error) {
	err = tx.Rollback()
	return
}
