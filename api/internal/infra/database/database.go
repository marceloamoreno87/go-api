package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/marceloamoreno/goapi/config"
)

func GetDBConn() (*pgx.Conn, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://"+config.Environment.DBUser+":"+config.Environment.DBPassword+"@"+config.Environment.DBHost+":"+config.Environment.DBPort+"/"+config.Environment.DBName)
	if err != nil {
		log.Fatal(err)
	}
	return conn, err
}
