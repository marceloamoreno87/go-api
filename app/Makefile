createmigration:
	migrate create -ext sql -dir sql/migrations -seq $(name) 
	chown -R 1000:1000 sql

migrate:
	migrate -path sql/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:5432/${DB_NAME}?sslmode=disable" up
	sqlc generate

migratedown:
	migrate -path sql/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:5432/${DB_NAME}?sslmode=disable" down
	sqlc generate

swaggen:
	swag init
	chown -R 1000:1000 docs
