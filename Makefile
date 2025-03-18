.PHONY: clean init db_migrate db_seed unit_test e2e_test bin/main all

POSTGRESQL_URL=$(shell awk -F ': ' '/migrateConnString:/ {sub(/migrateConnString: /, "", $$0); print $$0}' config.yaml)

all: clean init db_migrate db_seed unit_test e2e_test

clean:
	@docker-compose down --volumes

init:
	@go mod tidy
	@docker-compose up --build --no-start && docker-compose start
	@echo "Wait for database to be ready..."
	@sleep 10

db_migrate:
	@echo "Migrating database..."	
	@echo "y" | migrate -database ${POSTGRESQL_URL} -path database/postgres/migrations down
	@migrate -database ${POSTGRESQL_URL} -path database/postgres/migrations up

db_seed:
	@echo "Seeding database..."	
	@echo "y" | migrate -database ${POSTGRESQL_URL} -path database/postgres/seeds down
	@migrate -database ${POSTGRESQL_URL} -path database/postgres/seeds up

unit_test:
	@echo "Running unit tests..."
	@go test -v internal/adapter/repository/postgres/product/*_test.go

e2e_test:
	@echo "Running e2e tests..."
	@go run tests/e2e-test.go

bin/main: cmd/main.go
	@echo "Building binary..."
	@go build -o $@ $<