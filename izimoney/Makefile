createmigration:
	migrate create -ext sql -dir sql/migrations -seq $(name) && chown -R 1000:1000 sql
migrate:
	migrate -path sql/migrations -database "postgres://izimoney:izimoney@izimoney_db:5432/izimoney?sslmode=disable" up

migratedown:
	migrate -path sql/migrations -database "postgres://izimoney:izimoney@localhost:5432/izimoney_db?sslmode=disable" down