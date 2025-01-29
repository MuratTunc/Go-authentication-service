// constants.go
package main

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

// Database connection and service settings
const (
	SSLMode            = "disable"
	BackEndServicePort = ":8080"
	BackEndServiceName = "AUTHENTICATION-SERVICE"
	DBHost             = "localhost"
	DBUser             = "postgres"
	DBPassword         = "authentication-pass"
	DBName             = "authentication_db"
	DBPort             = "5439"
)
