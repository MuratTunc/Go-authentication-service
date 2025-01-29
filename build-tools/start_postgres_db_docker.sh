#!/bin/bash

# Check if Docker is installed
if ! command -v docker &> /dev/null
then
    echo "Docker is not installed. Please install Docker first."
    exit 1
fi

# Define container name and database credentials
CONTAINER_NAME="authentication-postgres"
DB_PASSWORD="authentication-pass"
DB_NAME="authentication_db"
DB_PORT="5439"
TABLE_NAME="users"

# Check if the port is already in use
if lsof -i :"${DB_PORT}" | grep LISTEN; then
    echo "Port ${DB_PORT} is already in use. Please choose a different port."
    exit 1
fi

# Check if the container name is already in use
if [ "$(docker ps -aq -f name=${CONTAINER_NAME})" ]; then
    echo "Container name '${CONTAINER_NAME}' is already in use."
    echo "Stopping and removing the existing container..."
    docker stop "${CONTAINER_NAME}"
    docker rm "${CONTAINER_NAME}"
fi

# Pull the latest PostgreSQL Docker image
echo "Pulling the latest PostgreSQL Docker image..."
docker pull postgres:latest

# Run the PostgreSQL container with the specified configuration
echo "Starting the PostgreSQL container..."
docker run -d \
  --name "${CONTAINER_NAME}" \
  -e POSTGRES_PASSWORD="${DB_PASSWORD}" \
  -e POSTGRES_DB="${DB_NAME}" \
  -p "${DB_PORT}":5432 \
  -v authentication-volume:/var/lib/postgresql/data \
  postgres:latest

# Wait for the database to be ready
echo "Waiting for PostgreSQL to start..."
sleep 15

# Create the authentication table in the database
echo "Creating the table '${TABLE_NAME}'..."

# Execute SQL commands to create the table
docker exec -it "${CONTAINER_NAME}" psql -U postgres -d "${DB_NAME}" -c "
CREATE TABLE IF NOT EXISTS ${TABLE_NAME} (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    mailaddress VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    activated BOOLEAN DEFAULT FALSE,
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    login_status BOOLEAN DEFAULT FALSE
);
"

# Confirm the table creation
echo "Table '${TABLE_NAME}' created successfully!"

# End of script
