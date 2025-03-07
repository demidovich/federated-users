.PHONY: vendor

help: ## This help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Docker
# ============================================================================

export HOST_UID ?= $(shell id -u)
export HOST_GID ?= $(shell id -g)

IMAGES := $(shell docker images -q "federated-*")
CONTAINERS := $(shell docker ps -aq --filter name=federated-)

up: ## Up application
	mkdir -p docker/var/postgres
	docker-compose -f docker-compose.yml up -d --remove-orphans

down: ## Down application
	docker stop $(CONTAINERS)

cleanup: down ## Remove containers
	docker rm -f --volumes $(CONTAINERS)

cleanup-full: cleanup ## Remove containers, images and networks
	docker rmi -f $(IMAGES)
	docker network rm -f federated-network

logs: ## Logs
	docker-compose logs --follow

# Postgres
# ============================================================================

migrate: ## Apply database migrations
	go run cmd/main.go migrate

migrate-rollback: ## Rollback database migrations
	go run cmd/main.go migrate:rollback

migrate-force: ## Apply force database migrations
	go run cmd/main.go migrate:force

migrate-status: ## Status database migrations
	go run cmd/main.go migrate:status

psql: ## Psql
	docker exec --interactive --tty federated-postgres psql federated_db

shell-postgres: ## Shell of postgresql container
	docker exec --interactive --tty federated-postgres /bin/bash

# Modules
# ============================================================================

vendor: ## Go mod vendor
	go mod tidy
	go mod vendor
