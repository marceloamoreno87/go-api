package configs

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/marceloamoreno/izimoney/internal/db"
)

func Queries() *db.Queries {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://izimoney:izimoney@db:5432/izimoney")
	if err != nil {
		log.Fatal(err)
	}
	return db.New(conn)
}

func CloseConnection(conn *pgx.Conn) {
	conn.Close(context.Background())
}
