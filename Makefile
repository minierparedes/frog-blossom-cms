postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root frog-blossom-db

dropdb:
	docker exec -it postgres12 dropdb frog-blossom-db

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/frog-blossom-db?sslmode=disable" -verbose up

migrateforce:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/frog-blossom-db?sslmode=disable" -verbose force 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/frog-blossom-db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run cmd/main.go

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server test