package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/marceloamoreno/izimoney/config"
	"github.com/marceloamoreno/izimoney/pkg/sqlc/db"
)

func GetQueries() *db.Queries {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://"+config.Environment.DBUser+":"+config.Environment.DBPassword+"@"+config.Environment.DBHost+":5432/"+config.Environment.DBName)
	if err != nil {
		log.Fatal(err)
	}
	return db.New(conn)
}
