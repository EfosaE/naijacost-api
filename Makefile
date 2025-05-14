.PHONY: migrate-create migrate-up migrate-down sqlc

# Create a new migration file
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir internal/db/migrations -seq $$name

migrate-up:
	@if [ -z "$$PROD_DB_URL" ]; then \
		echo "PROD_DB_URL is not set. Please export it first."; \
		exit 1; \
	fi && \
	migrate -database "$$PROD_DB_URL" -path ./internal/db/migrations up



# Rollback the most recent migration
migrate-down:
	. .env && migrate -database "$$(grep ^PROD_DB_URL .env | cut -d '=' -f2-)" -path ./internal/db/migrations down 1

# Generate sqlc code
sqlc:
	sqlc generate

# run the server
start:
	go run cmd/server/main.go