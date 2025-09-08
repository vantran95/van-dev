ifndef PROJECT_NAME
PROJECT_NAME := restful
endif

ifndef PRODUCTION_ENVIRONMENT:
PRODUCTION_ENVIRONMENT := prod
endif

ifndef DOCKER_BIN:
DOCKER_BIN := docker
endif

ifndef DOCKER_COMPOSE_BIN:
DOCKER_COMPOSE_BIN := docker-compose
endif

COMPOSE := ${DOCKER_COMPOSE_BIN} -f docker-compose.yaml
ifdef CONTAINER_SUFFIX
COMPOSE := ${COMPOSE} -f docker-compose.ci.yaml -p ${CONTAINER_SUFFIX}
else
COMPOSE := ${COMPOSE} -f docker-compose.yaml
endif

# ----------------------------
# api Methods
# ----------------------------
API_COMPOSE = ${COMPOSE} run --name api-service --rm --service-ports

# Environment variables
#DB_URL=postgres://postgres:postgres@localhost:5432/restful?sslmode=disable
MIGRATE_CMD=docker-compose run --rm -v $(PWD)/migrations:/migrations pg-migrate

DB_URL=postgres://postgres:@localhost:5432/restful?sslmode=disable


api-gen-models:
	${API_COMPOSE} sh -c 'sqlboiler --wipe psql && goimports -w internal/repository/orm/*.go'

# Run Migrations Up (Apply all pending migrations)
api-migrate-up:
	@echo "Applying all up migrations..."
	$(MIGRATE_CMD) -path=/migrations -database "$(DB_URL)" up

# Run Migrations Down (Rollback last migration)
api-migrate-down:
	@echo "Rolling back the last migration..."
	$(MIGRATE_CMD) -path=/migrations -database "$(DB_URL)" down 1

api-pg-migrate:
	${COMPOSE} run --rm pg-migrate sh -c './migrate -path ./migrations -database $$PG_URL up'
api-pg-drop:
	${COMPOSE} run --rm pg-migrate sh -c './migrate -path ./migrations -database $$PG_URL drop'
api-pg-redo: api-pg-drop api-pg-migrate


