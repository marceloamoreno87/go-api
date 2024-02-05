package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/marceloamoreno/goapi/config"
)

func GetDBConn() (*sql.DB, error) {
	databaseConfig := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Environment.DBHost, config.Environment.DBPort, config.Environment.DBUser, config.Environment.DBPassword, config.Environment.DBName)
	db, err := sql.Open("postgres", databaseConfig)

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	return db, err
}

func GetDbMockConn() (*sql.DB, error) {
	databaseConfig := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Environment.DBHost, config.Environment.DBPort, config.Environment.DBUser, config.Environment.DBPassword, config.Environment.DBMockName)
	db, err := sql.Open("postgres", databaseConfig)

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	return db, err
}
