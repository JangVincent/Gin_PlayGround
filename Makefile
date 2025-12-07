include .env
export

.PHONY: migrate-up migrate-down migrate-create migrate-force migrate-version sqlc-generate dev build clean

migrate-up:
	migrate -database "$(DATABASE_URL)" -path database/migrations up

migrate-down:
	migrate -database "$(DATABASE_URL)" -path database/migrations down

migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir database/migrations -seq $$name

migrate-force:
	@read -p "Enter version: " version; \
	migrate -database "$(DATABASE_URL)" -path database/migrations force $$version

migrate-version:
	migrate -database "$(DATABASE_URL)" -path database/migrations version

db-gen:
	sqlc generate

dev:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

clean:
	rm -rf bin/
	rm -rf internal/db/

test:
	go test -v ./...
