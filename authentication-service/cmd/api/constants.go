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
}
