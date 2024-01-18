package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/marceloamoreno/izimoney/configs"
	"github.com/marceloamoreno/izimoney/pkg/sqlc/db"
)

func Db() *db.Queries {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://"+configs.Environment.DBUser+":"+configs.Environment.DBPassword+"@"+configs.Environment.DBHost+":5432/"+configs.Environment.DBName)
	if err != nil {
		log.Fatal(err)
	}
	return db.New(conn)
}
