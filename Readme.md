Authentication Micro-service via PostgreSQL and Golang-GORM-Gokit

In this project, I will explain the communication between the PostgreSQL database running on Docker and the authentication go service running on Docker using the Golang SW language.![AuthMain](https://github.com/user-attachments/assets/bdc7fa8c-8926-4d15-a746-110d5dab74e2)

— Project Structure —
— Project Folder structure —

1-Environment file:

If you want your app to automatically read .env variables, install github.com/joho/godotenv:

go get github.com/joho/godotenv
— environment file —

We will use these parameters in docker compose yaml and golgo get github.com/joho/godotenvang code.

2-Docker file:
— docker file —

3-Docker Compose File:
— docker compose yaml —

4-MakeFile:

# Environment file and variables
ENV_FILE := .env

## up_build: Stops and removes all running containers, builds the project and starts docker-compose
build: stop_all_containers down build_authentication_service
 echo "🚀 Building (when required) and starting docker images with environment variables..."
 docker-compose up --build -d
 echo "✅ Docker images built and started!"
 echo "📜 Fetching logs for all services..."
 docker-compose logs --tail=50
 echo "🚀 Runnig Containers:"
 docker ps

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

— Generating Services Hierarchy —

Start make file:

sudo make -s build

mutu@mutu:~/projects/Go-authentication-service/build-tools$ sudo make -s build
[sudo] password for mutu: 
🔍 Checking for running containers...
🛑 Stopping all running Docker containers...
c5a72fc90cc8
ca2bfc33b87a
🗑️ Removing all stopped containers...
c5a72fc90cc8
ca2bfc33b87a
✅ All containers stopped and removed.
🛑 Stopping docker-compose...
Removing network build-tools_default
✅ Done!
🚀 Building authentication-service binary...
✅ Done!
🚀 Building (when required) and starting docker images with environment variables...
Creating network "build-tools_default" with the default driver
Building auth-service
[+] Building 18.4s (15/15) FINISHED                                                                                       docker:default
 => [internal] load build definition from authentication-service.dockerfile                                                         0.0s
 => => transferring dockerfile: 786B                                                                                                0.0s
 => [internal] load metadata for docker.io/library/alpine:latest                                                                    1.5s
 => [internal] load metadata for docker.io/library/golang:1.23-alpine                                                               1.5s
 => [internal] load .dockerignore                                                                                                   0.0s
 => => transferring context: 2B                                                                                                     0.0s
 => [builder 1/6] FROM docker.io/library/golang:1.23-alpine@sha256:47d337594bd9e667d35514b241569f95fb6d95727c24b19468813d596d5ae59  0.0s
 => [stage-1 1/3] FROM docker.io/library/alpine:latest@sha256:56fa17d2a7e7f168a043a2712e63aed1f8543aeafdcee47c58dcffe38ed51099      0.0s
 => [internal] load build context                                                                                                   0.0s
 => => transferring context: 5.68kB                                                                                                 0.0s
 => CACHED [builder 2/6] WORKDIR /app                                                                                               0.0s
 => [builder 3/6] COPY go.mod go.sum ./                                                                                             0.1s
 => [builder 4/6] RUN go mod download                                                                                               8.4s
 => [builder 5/6] COPY cmd/api /app/cmd/api                                                                                         0.1s
 => [builder 6/6] RUN go build -o authenticationServiceApp ./cmd/api                                                                8.0s
 => CACHED [stage-1 2/3] WORKDIR /app                                                                                               0.0s
 => [stage-1 3/3] COPY --from=builder /app/authenticationServiceApp .                                                               0.1s
 => exporting to image                                                                                                              0.1s
 => => exporting layers                                                                                                             0.1s
 => => writing image sha256:e1d22c3e54eeb655d48583de91ccea233ab028d42fca704547de749917b0a6c3                                        0.0s
 => => naming to docker.io/library/authentication-service-img                                                                       0.0s
Creating auth-db ... done
Creating build-tools_auth-service_1 ... done
✅ Docker images built and started!
📜 Fetching logs for all services...
Attaching to build-tools_auth-service_1, auth-db
auth-service_1  | 🔧 Loaded Environment Variables:
auth-service_1  | DBUser: auth_user
auth-service_1  | DBPassword: auth_password
auth-service_1  | DBName: auth_db
auth-service_1  | DBPort: 5432
auth-service_1  | DBHost: auth-db
auth-service_1  | ServicePort: 8080
auth-service_1  | ServiceName: AUTHENTICATION-SERVICE
auth-service_1  | DATABASE connection success!
auth-db         | 
auth-db         | PostgreSQL Database directory appears to contain a database; Skipping initialization
auth-db         | 
auth-db         | 2025-02-01 19:07:16.787 UTC [1] LOG:  starting PostgreSQL 15.10 (Debian 15.10-1.pgdg120+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 12.2.0-14) 12.2.0, 64-bit
auth-db         | 2025-02-01 19:07:16.787 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
auth-db         | 2025-02-01 19:07:16.787 UTC [1] LOG:  listening on IPv6 address "::", port 5432
auth-db         | 2025-02-01 19:07:16.795 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
auth-db         | 2025-02-01 19:07:16.809 UTC [29] LOG:  database system was shut down at 2025-02-01 19:06:55 UTC
auth-db         | 2025-02-01 19:07:16.816 UTC [1] LOG:  database system is ready to accept connections
🚀 Runnig Containers:
CONTAINER ID   IMAGE                        COMMAND                  CREATED        STATUS                  PORTS                                       NAMES
933759e0b769   authentication-service-img   "/app/authentication…"   1 second ago   Up Less than a second   0.0.0.0:8080->8080/tcp, :::8080->8080/tcp   build-tools_auth-service_1
e57f21a15aae   postgres:15                  "docker-entrypoint.s…"   1 second ago   Up 1 second             0.0.0.0:5432->5432/tcp, :::5432->5432/tcp   auth-db

We have 2 docker containers, the first one is our authentication service listening on port 8080, the second one is our PostgreSQL database service listening on port 5432.

Test with “HTTP” post request from Ubuntu terminal.
Step 1: Register User (Create a New User)

You can use curl to send a POST request with user information in the body. Here's the command:
— Register User HTTP post request —
Step 2: Login User (Verify Credentials)

After registering, you can send a POST request to the /login endpoint to verify the user's credentials:
— Login User HTTP post request —

You can check the PostgreSQL database table results from the terminal by following these steps:
Access the PostgreSQL Container

Since you are running PostgreSQL in Docker (auth-db service), first enter the PostgreSQL container:

If you notice, we do not write the user password directly in the database, we save it by calculating the hash code. You can examine the encryption function in the handlers.go file

// HashPassword hashes a password using bcrypt
func (app *Config) HashPassword(password string) (string, error) {
 hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
 if err != nil {
  return "", err
 }
 return string(hashedPassword), nil
}

Micro-service architecture: The server is started in the main.go file, request acceptance definitions are made for the necessary handle functions in the routes.go file and CORS settings are handled. Function bodies for 2 request responses are created in the handlers.go file.

main.go:

The entry point for our authentication service. It does the following:

    Loads environment variables
    Connects to the PostgreSQL database
    Auto-migrates the User model
    Starts the HTTP server

routes.go:

Defines and configures the API routes for our authentication service. It uses the Chi router to handle HTTP requests and middleware to add functionality like logging, error recovery, and CORS support.

handlers.go:

Defines the user authentication handlers for our Golang authentication service. It includes functions for user registration and login while handling password hashing and database interactions.

GORM and db table structure:

The User struct is our GORM model, which defines how the users table is structured in your PostgreSQL database.

// User model for GORM
type User struct {
 ID          uint      `gorm:"primaryKey"`
 Username    string    `gorm:"unique;not null"`
 MailAddress string    `gorm:"unique;not null"`
 Password    string    `gorm:"not null"`
 Activated   bool      `gorm:"default:false"`
 LoginStatus bool      `gorm:"default:false"`
 CreatedAt   time.Time `gorm:"autoCreateTime"`
 UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

Each field in the struct corresponds to a column in the users table.
GORM uses this struct to automatically create or migrate the table.

Enforces Constraints (Data Integrity)

    gorm:"primaryKey" → Marks ID as the primary key.
    gorm:"unique;not null" → Ensures Username & MailAddress must be unique and cannot be NULL.
    gorm:"not null" → Ensures Password is always required.
    gorm:"default:false" → Activated and LoginStatus default to false if not provided.
    gorm:"autoCreateTime" → Automatically sets CreatedAt when a new record is created.
    gorm:"autoUpdateTime" → Updates UpdatedAt whenever the record is modified.

This ensures data consistency and prevents duplicate users in your database.
Enables Auto-Migration

If you call db.AutoMigrate(&User{}), GORM will:

    Create the users table if it doesn’t exist.
    Add new columns if the struct changes.
    Keep your database schema in sync with your Go code.

Source Code:

HTTP:https://github.com/MuratTunc/Go-authentication-service.git

SSH:git@github.com:MuratTunc/Go-authentication-service.git

In this study, I tried to introduce the simple authentication micro service created with the go language and the PostgreSQL database on docker as much as I could. You can use this source code in web applications or any cloud authentication work.

Thank you for your time.
Murat Tunc
Written by Murat Tunc
0 Followers
·
3 Following
Edit profile
No responses yet
Murat Tunc
Murat Tunc




