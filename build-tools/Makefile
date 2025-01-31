# Environment file and variables
ENV_FILE := .env

## Ensure environment file exists and GOFULLPATH is set
setup_env:
	@if [ ! -f $(ENV_FILE) ]; then \
		echo "❌ Error: $(ENV_FILE) not found. Please create it first."; \
		exit 1; \
	fi; \
	set -a; . ./$(ENV_FILE); set +a; \
	if [ -z "$$GOFULLPATH" ]; then \
		echo "❌ Error: GOFULLPATH is not set in .env"; \
		exit 1; \
	fi; \
	echo "✅ Environment variables loaded successfully!"

## up_build: Stops and removes all running containers, builds the project, checks ports, and starts docker-compose
up_build: setup_env stop_all_containers build_authentication_service
	echo "Checking if service ports are available..."
	$(MAKE) check_ports
	echo "🚀 Building (when required) and starting docker images with environment variables..."
	docker-compose up --build -d
	echo "✅ Docker images built and started!"
	echo "📜 Fetching logs for all services..."
	docker-compose logs --tail=50
	echo "🧹 Cleaning up built binaries..."
	$(MAKE) clean
	echo "✅ up_build process completed!"

## stop_all_containers: Stops and removes all running Docker containers (if any exist)
stop_all_containers:
	echo "🔍 Checking for running containers..."
	@if [ -n "$$(docker ps -q)" ]; then \
		echo "🛑 Stopping all running Docker containers..."; \
		docker stop $$(docker ps -q); \
		echo "🗑️ Removing all stopped containers..."; \
		docker rm $$(docker ps -aq); \
		echo "✅ All containers stopped and removed."; \
	else \
		echo "⚡ No running containers found. Skipping stop and remove."; \
	fi

## check_ports: Checks if service ports are available
check_ports:
	echo "Checking if service ports are available..."
	if lsof -i :$$AUTHENTICATION_SERVICE_PORT; then \
		echo "❌ Error: Port $$AUTHENTICATION_SERVICE_PORT is already in use."; \
		exit 1; \
	else \
		echo "✅ Ports are available."; \
	fi

## down: stop docker-compose
down:
	echo "🛑 Stopping docker-compose..."
	docker-compose down
	echo "✅ Done!"

## build_authentication_service: Builds the authentication-service binary as a Linux executable
build_authentication_service:
	echo "🚀 Building authentication-service binary..."
	@set -a; . ./$(ENV_FILE); set +a; \
	cd ../authentication-service && $$GOFULLPATH build -o $$AUTHENTICATION_SERVICE_BINARY ./cmd/api
	echo "✅ Done!"

## logs: Shows logs from all services
logs:
	echo "📜 Fetching last 50 logs for all services..."
	docker-compose logs --tail=50 -f

## help: Displays the list of available commands
help:
	@grep -E '^##' $(MAKEFILE_LIST) | sed -e 's/## //'

## clean: removes built binaries
clean:
	@echo "Cleaning up built binaries..."
	@set -a; . ./$(ENV_FILE); set +a; \
		echo "AUTHENTICATION_SERVICE_BINARY: $$AUTHENTICATION_SERVICE_BINARY"; \
		[ -f ../authentication-service/$$AUTHENTICATION_SERVICE_BINARY ] && rm -f ../authentication-service/$$AUTHENTICATION_SERVICE_BINARY || echo "File $$AUTHENTICATION_SERVICE_BINARY does not exist";
	@echo "Done!"