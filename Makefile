build:
	go build -o bin/api ./cmd/api

run:
	go run ./cmd/api

migrate-up:
	@go run ./cmd/migrate up

migrate-down:
	@go run ./cmd/migrate down

migrate-force:
	@go run ./cmd/migrate force $(v)
