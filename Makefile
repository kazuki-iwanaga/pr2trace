include .env

.PHONY: e2e
e2e:
	docker compose exec app bash -c "go run main.go --owner kazuki-iwanaga --repo pr2otel --number 7"

.PHONY: test
test:
	docker compose exec app bash -c "go test -v ./..."

.PHONY: fmt
fmt:
	docker compose exec app bash -c "go fmt ./..."

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

.PHONY: o2
o2:
	open http://localhost:$(O2_HTTP_PORT)
