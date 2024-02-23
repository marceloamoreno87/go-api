package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/marceloamoreno/goapi/config"
)

type DatabaseInterface interface {
	SetDbConn() (err error)
	GetDbConn() (db *sql.DB)
}

type Database struct {
	dbConn   *sql.DB
	driver   string
	host     string
	port     string
	user     string
	password string
	dbname   string
	sslmode  string
}

func NewDatabase() *Database {
	return &Database{
		driver:   "postgres",
		host:     config.NewEnv().GetDBHost(),
		port:     config.NewEnv().GetDBPort(),
		user:     config.NewEnv().GetDBUser(),
		password: config.NewEnv().GetDBPassword(),
		dbname:   config.NewEnv().GetDBName(),
		sslmode:  "disable",
	}
}

func (d *Database) SetDbConn() (err error) {
	databaseConfig := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", d.host, d.port, d.user, d.password, d.dbname, d.sslmode)
	d.dbConn, err = sql.Open(d.driver, databaseConfig)
	if err != nil {
		return
	}
	err = d.dbConn.Ping()
	return
}

func (d *Database) GetDbConn() (db *sql.DB) {
	return d.dbConn
}
