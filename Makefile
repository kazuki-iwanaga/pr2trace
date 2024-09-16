include .env

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: in
in:
	docker compose exec app bash

.PHONY: build
build:
	docker compose build --no-cache

.PHONY: ps
ps:
	docker compose ps -a

.PHONY: logs
logs:
	docker compose logs

.PHONY: logsf
logsf:
	docker compose logs -f
