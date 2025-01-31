package main

import (
	"fmt"
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
	DBUser     = getEnv("POSTGRES_DB_USER", "x")
	DBPassword = getEnv("POSTGRES_DB_PASSWORD", "x")
	DBName     = getEnv("POSTGRES_DB_NAME", "x")
	DBPort     = getEnv("POSTGRES_DB_PORT", "5432")
	DBHost     = getEnv("POSTGRES_DB_HOST", "x")
	ServerPort = getEnv("AUTHENTICATION_SERVICE_PORT", "x")
)

// getEnv reads an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// PrintEnvVariables prints all environment variables for debugging
func PrintEnvVariables() {
	fmt.Println("🔧 Loaded Environment Variables:")
	fmt.Printf("DBUser: %s\n", DBUser)
	fmt.Printf("DBPassword: %s\n", DBPassword)
	fmt.Printf("DBName: %s\n", DBName)
	fmt.Printf("DBPort: %s\n", DBPort)
	fmt.Printf("DBHost: %s\n", DBHost)
	fmt.Printf("ServerPort: %s\n", ServerPort)
}
