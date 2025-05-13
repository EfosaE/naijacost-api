.PHONY: migrate-create migrate-up migrate-down sqlc

# Create a new migration file
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir internal/db/migrations -seq $$name

# Run all pending migrations
migrate-up:
	migrate -database "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable" -path ./internal/db/migrations up

# Rollback the most recent migration
migrate-down:
	migrate -database "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable" -path ./internal/db/migrations down 1

# Generate sqlc code
sqlc:
	sqlc generate

# run the server
start:
	go run cmd/server/main.go