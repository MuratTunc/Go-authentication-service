// constants.go
package main

import (
	"os"
)

// Constants for error and success messages
const (
	ErrInvalidRequestBody = "Invalid request body"
	ErrHashingPassword    = "Error hashing password"
	ErrInsertingUser      = "Error inserting user"
	ErrUserNotFound       = "User not found"
	ErrInvalidCredentials = "Invalid credentials"
	UserCreatedSuccess    = "User created successfully"
	LoginSuccess          = "Login successful"
)

// Constants for environment variables
var (
	DBUser     = getEnv("POSTGRES_DB_USER", "postgres")
	DBPassword = getEnv("POSTGRES_DB_PASSWORD", "yourpassword")
	DBName     = getEnv("POSTGRES_DB_NAME", "authentication_db")
	DBPort     = getEnv("POSTGRES_DB_PORT", "5439")
	DBHost     = getEnv("POSTGRES_DB_HOST", "localhost")
	DBSSLMode  = getEnv("POSTGRES_DBSSLMode", "disabled")
	ServerPort = getEnv("AUTHENTICATION_SERVICE_PORT", "8082")
)

// getEnv reads an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
