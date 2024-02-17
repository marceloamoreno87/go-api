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

func (d *DefaultRepository) BeginTx(ctx context.Context, options *sql.TxOptions) (tx *sql.Tx, err error) {
	tx, err = d.DbConn.BeginTx(ctx, options)
	return
}

func (d *DefaultRepository) CommitTx(tx *sql.Tx) (err error) {
	err = tx.Commit()
	return
}

func (d *DefaultRepository) RollbackTx(tx *sql.Tx) (err error) {
	err = tx.Rollback()
	return
}
