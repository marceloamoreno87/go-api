package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DatabaseMock struct {
	dbConn *sql.DB
}

var DbMock DatabaseInterface

func NewDatabaseMock(db *sql.DB) {
	dbMock := &DatabaseMock{
		dbConn: db,
	}
	DbMock = dbMock
}

func (d *DatabaseMock) SetDbConn() (err error) {
	return
}

func (d *DatabaseMock) GetDbConn() (db *sql.DB) {
	return d.dbConn
}
