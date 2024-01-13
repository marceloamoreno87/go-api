createmigration:
	migrate create -ext sql -dir=sql/migrations -seq $(name)

migrate:
	migrate -path=sql/migrations -database "postgres://izimoney:izimoney@localhost:5432/izimoney?sslmode=disable" up

migratedown:
	migrate -path=sql/migrations -database "postgres://izimoney:izimoney@localhost:5432/izimoney?sslmode=disable" down