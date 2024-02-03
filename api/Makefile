createmigration:
	migrate create -ext sql -dir pkg/sqlc/migrations -seq $(name) 
	chown -R 1000:1000 sql

migrate:
	migrate -path internal/sql/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up
	sqlc generate

migratedown:
	migrate -path internal/sql/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" down
	sqlc generate

swaggen:
	swag init --parseDependency
	chown -R 1000:1000 docs

testgo:
	migrate -path internal/sql/migrations -database "postgres://${DB_USER_MOCK}:${DB_PASSWORD_MOCK}@${DB_HOST}:${DB_PORT}/${DB_NAME_MOCK}?sslmode=disable" up
	go test -v ./...

build go:
	go test -v ./...
	go build