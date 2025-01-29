AUTHENTICATION_SERVICE_BINARY=authenticationServiceApp
GOFULLPATH := /usr/local/go/bin/go

## up_build: stops docker-compose (if running), builds all projects, starts docker compose, and runs PostgreSQL DB setup
up_build: build_authentication_service_binary build_postgres_db
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_back_end_service: builds the build_back_end_service binary as a linux executable
build_authentication_service_binary:
	@echo "Building binary..."
	cd ../authentication-service && env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOFULLPATH) build -o ${AUTHENTICATION_SERVICE_BINARY} ./cmd/api
	@echo "Done!"

## build_postgres_db: runs the PostgreSQL setup script
build_postgres_db:
	@echo "Running PostgreSQL setup script..."
	./start_postgres_db_docker.sh
	@echo "PostgreSQL setup complete!"