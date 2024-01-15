package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/marceloamoreno/izimoney/pkg/sqlc/db"
)

func Db() *db.Queries {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://"+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@"+os.Getenv("DB_HOST")+":5432/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	return db.New(conn)
}
