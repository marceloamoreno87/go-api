createmigration:
	migrate create -ext sql -dir internal/sql/migrations -seq $(name) 
	chown -R 1000:1000 internal/sql/migrations

sqlcgen:
	sqlc generate
	chown -R 1000:1000 internal/shared/db/

migrate:
	rm -rf internal/shared/db/
	migrate -path internal/sql/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up
	sqlc generate
	chown -R 1000:1000 pkg/sqlc

migratedown:
	migrate -path internal/sql/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" down

swaggen:
	swag init --parseDependency
	chown -R 1000:1000 docs

testgo:
	go test -v ./...

build go:
	go test -v ./...
	go build