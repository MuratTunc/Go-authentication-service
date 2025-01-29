#!/bin/bash

# Define colors
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Check if Docker is installed
if ! command -v docker &> /dev/null
then
    echo -e "${RED}Docker is not installed. Please install Docker first.${NC}"
    exit 1
fi

# Define container name and database credentials
CONTAINER_NAME="authentication-postgres"
DB_PASSWORD="authentication-pass"
DB_NAME="authentication_db"
DB_PORT="5439"
TABLE_NAME="users"

# Check if the container name is already in use
if [ "$(docker ps -aq -f name=${CONTAINER_NAME})" ]; then
    echo "Container name '${CONTAINER_NAME}' is already in use."
    echo "Stopping and removing the existing container..."
    docker stop "${CONTAINER_NAME}" && docker rm "${CONTAINER_NAME}"
    if [ $? -ne 0 ]; then
        echo -e "${RED}Failed to stop/remove existing container!${NC}"
        exit 1
    fi
fi

# Check if the port is already in use
if lsof -i :"${DB_PORT}" | grep LISTEN; then
    echo -e "${RED}Port ${DB_PORT} is already in use. Please choose a different port.${NC}"
    exit 1
fi

# Pull the latest PostgreSQL Docker image
echo "Pulling the latest PostgreSQL Docker image..."
docker pull postgres:latest
if [ $? -ne 0 ]; then
    echo -e "${RED}Failed to pull PostgreSQL image!${NC}"
    exit 1
fi

# Run the PostgreSQL container with the specified configuration
echo "Starting the PostgreSQL container..."
docker run -d \
  --name "${CONTAINER_NAME}" \
  -e POSTGRES_PASSWORD="${DB_PASSWORD}" \
  -e POSTGRES_DB="${DB_NAME}" \
  -p "${DB_PORT}":5432 \
  -v authentication-volume:/var/lib/postgresql/data \
  postgres:latest

if [ $? -ne 0 ]; then
    echo -e "${RED}Failed to start PostgreSQL container!${NC}"
    exit 1
fi

# Wait for the database to be ready
echo "Waiting for PostgreSQL to start..."
sleep 15

# Check if the container is running
if [ ! "$(docker ps -q -f name=${CONTAINER_NAME})" ]; then
    echo -e "${RED}Error: PostgreSQL container failed to start.${NC}"
    exit 1
fi

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
    updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
"

# Add a trigger function to update `updated_time`
docker exec -it "${CONTAINER_NAME}" psql -U postgres -d "${DB_NAME}" -c "
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS \$\$
BEGIN
    NEW.updated_time = NOW();
    RETURN NEW;
END;
\$\$ LANGUAGE plpgsql;
"

# Create the trigger to update `updated_time` on any row update
docker exec -it "${CONTAINER_NAME}" psql -U postgres -d "${DB_NAME}" -c "
CREATE TRIGGER trigger_update_timestamp
BEFORE UPDATE ON ${TABLE_NAME}
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
"

echo -e "${GREEN}Table '${TABLE_NAME}' created successfully with automatic updated_time handling!${NC}"
echo -e "${GREEN}Table '${TABLE_NAME}' created successfully!${NC}"

# End of script
