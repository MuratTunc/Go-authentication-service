# Environment file and variables
ENV_FILE := .env

## up_build: Stops and removes all running containers, builds the project, checks ports, and starts docker-compose
up_build: stop_all_containers clean build_authentication_service
	@echo "Checking if service ports are available..."
	$(MAKE) check_ports
	@echo "Building (when required) and starting docker images with environment variables..."
	docker-compose up --build -d
	@echo "Docker images built and started!"
	@echo "Fetching logs for all services..."
	docker-compose logs --tail=50
	@echo "Cleaning up built binaries..."
	$(MAKE) clean
	@echo "✅ up_build process completed!"

## stop_all_containers: Stops and removes all running Docker containers
stop_all_containers:
	@echo "Stopping all running Docker containers..."
	-docker stop $$(docker ps -q)
	@echo "Removing all stopped containers..."
	-docker rm $$(docker ps -aq)
	@echo "✅ All containers stopped and removed."

## check_ports: Checks if service ports are available
check_ports:
	@echo "Checking if service ports are available..."
	@set -a; . ./$(ENV_FILE); set +a; \
		echo "Authentication Service Binary: $$AUTHENTICATION_SERVICE_BINARY"; \
		if lsof -i :$$AUTHENTICATION_SERVICE_PORT; then \
			echo "❌ Error: One or more ports are already in use."; \
			exit 1; \
		else \
			echo "✅ Ports are available."; \
		fi

## down: stop docker-compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "✅ Done!"

## clean: Removes built binaries
clean:
	@echo "Cleaning up built binaries..."
	@set -a; . ./$(ENV_FILE); set +a; \
		echo "AUTHENTICATION_SERVICE_BINARY: $$AUTHENTICATION_SERVICE_BINARY"; \
		[ -f ../authentication-service/$$AUTHENTICATION_SERVICE_BINARY ] && rm -f ../authentication-service/$$AUTHENTICATION_SERVICE_BINARY || echo "ℹ️ File $$AUTHENTICATION_SERVICE_BINARY does not exist";
	@echo "✅ Done!"

## build_authentication_service: Builds the authentication-service binary as a Linux executable
build_authentication_service:
	@echo "Building authentication-service binary..."
	@set -a; . ./$(ENV_FILE); set +a; \
		cd ../authentication-service && env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $$GOFULLPATH build -o $$AUTHENTICATION_SERVICE_BINARY ./cmd/api
	@echo "✅ Done!"

## logs: Shows logs from all services
logs:
	@echo "Fetching logs for all services..."
	docker-compose logs -f

## help: Displays the list of available commands
help:
	@grep -E '^##' $(MAKEFILE_LIST) | sed -e 's/## //'
