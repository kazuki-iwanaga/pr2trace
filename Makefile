include .env

#===============================================================================
# docker compose
#===============================================================================
.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: ps
ps:
	docker compose ps -a

.PHONY: logs
logs:
	docker compose logs

.PHONY: logsf
logsf:
	docker compose logs -f
#===============================================================================
