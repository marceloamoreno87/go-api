package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/marceloamoreno/izimoney/pkg/sqlc/db"
)

func Db() *db.Queries {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://izimoney:izimoney@db:5432/izimoney")
	if err != nil {
		log.Fatal(err)
	}
	return db.New(conn)
}
