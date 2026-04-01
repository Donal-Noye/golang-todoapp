include .env
export

#export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d todoapp-postgres

env-down:
	@docker compose down todoapp-postgres

env-cleanup:
	@bash env-cleanup.sh

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

migrate-create:
	@docker compose run --rm todoapp-postgres-migrate create -ext sql -dir /migrations -seq $(seq)

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@docker compose run --rm todoapp-postgres-migrate \
		-source file:///migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		$(action)

migrate-force:
	docker compose run --rm todoapp-postgres-migrate \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		force 0

db-reset:
	docker compose down -v
	docker compose up -d todoapp-postgres

todoapp-run:
	@export LOGGER_FOLDER=./out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run cmd/todoapp/main.go