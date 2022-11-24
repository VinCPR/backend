DB_URL?=postgresql://root:secret@localhost:5432/softcon?sslmode=disable

test:
	@PROJECT_PATH=$(shell pwd) go test -cover ./...

postgres:
	docker run --name softcon_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.1-alpine

createdb:
	docker exec -it softcon_postgres createdb --username=root --owner=root softcon

dropdb:
	docker exec -it softcon_postgres dropdb softcon

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

server:
	go run cmd/main.go

gen-swagger:
	swag init --parseDependency --parseInternal -g ./cmd/main.go

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server gen-swagger

#remove-infras:
#	docker-compose stop
#	docker-compose rm -f
#
#init-infras: remove-infras
#	docker-compose up -d $(DB_CONTAINER)
#	@echo "Waiting for database connection..."
#	@while ! docker exec $(DB_CONTAINER) pg_isready -h localhost -p 5432 > /dev/null; do \
#		sleep 1; \
#	done
