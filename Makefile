ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "v1consortium"
DOCKER_NAME = "v1consortium"

include ./hack/hack-cli.mk
include ./hack/hack.mk


.PHONY: setup
setup:
	@# Shut down and clear out any local postgres state
	docker compose stop postgres
	rm -rf .local/postgres
	@# Install the required ssl certs
	./bin/create-localhost-certs
	@# Start the database docker container
	docker compose up -d --wait postgres
	@# Wait for the database to be ready
	@until PGPASSWORD=password psql "postgres://postgres:password@localhost:5432?sslmode=disable" -c "SELECT 1" >/dev/null 2>&1; do \
		echo "PostgreSQL is unavailable - retrying..."; \
		sleep 2; \
	done
# 	@# Run database migrations
# 	make migrate up
# 	go install github.com/riverqueue/river/cmd/river@latest
# 	river migrate-up --database-url 'postgres://postgres:password@localhost?sslmode=disable'
# 	@# Seed the database
# 	psql "postgres://postgres:password@localhost:5432?sslmode=disable" -f .local/db/seed.sql
	@# Stop the docker containers
	docker compose stop postgres

.PHONY: dev
dev:
	docker compose up --build --watch