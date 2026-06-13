-include .env
export

export PROJECT_ROOT=$(shell pwd)
export LOGGER_FOLDER=${PROJECT_ROOT}/out/log

env-up:
	@docker compose up -d db-service-product

env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опасность утери данных. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down data-base forwarder-port && \
		rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi

logs-clean:
	@read -p "Очистить все логи? Опасность утери данных. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		rm -rf ${PROJECT_ROOT}/out/logs && \
		echo "логи очищены"; \
	else \
		echo "Очистка логов отменена"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Dont have seq"; \
		exit 1; \
	fi; \
		docker compose run --rm migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Dont have action"; \
		exit 1; \
	fi; \
	docker compose run --rm migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db-service-order:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

clean_migrate:
	@docker compose run --rm migrate \
		-path /migrations \
     	-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@data-base:5432/${POSTGRES_DB}?sslmode=disable force 0

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

port-forwarder-postgres-start:
	@docker compose up -d forwarder-port

port-forwarder-postgres-stop:
	@docker compose down forwarder-port

app-run:
	@go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/productapp/main.go