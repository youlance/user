DB_URL=postgresql://root:secret@localhost:5432/youlance_users?sslmode=disable

tidy:
	go mod tidy

migrate_create:
	migrate create -ext sql -dir migration -seq init_schema

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root youlance_users

dropdb:
	docker exec -it postgres dropdb youlance_users

migrateup:
	migrate -path migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path migrations -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate -f db/sqlc_config/sqlc.yaml

server:
	go run main.go

build-image:
	docker build -t userservice-image -f Dockerfile .

run-container:
	docker run --net=host --rm -d --name userservice userservice-image

.PHONY: tidy migrate_create postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc server