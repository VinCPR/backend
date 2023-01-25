DB_URL?=postgresql://root:secret@localhost:5432/vincpr?sslmode=disable

run-all:
	docker compose up

network:
	docker network create vincpr_network

postgres:
	docker run --name vincpr_postgres --network vincpr_network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.1-alpine

createdb:
	docker exec -it vincpr_postgres createdb --username=root --owner=root vincpr

dropdb:
	docker exec -it vincpr_postgres dropdb vincpr

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

test:
	go test -cover ./...

server:
	go run cmd/main.go

seed:
	go run seeder/main.go

sqlc:
	sqlc generate

gen-swagger:
	swag init --parseDependency --parseInternal -g ./cmd/main.go

.PHONY: run-all network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server gen-swagger