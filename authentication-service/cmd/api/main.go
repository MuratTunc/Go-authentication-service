// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	// Add any fields that you need here for your application configuration
	// For example, you might store your database connection or other configurations here.
	DB *gorm.DB
}

// connectToDB function to connect to PostgreSQL using constants
func connectToDB() (*gorm.DB, error) {
	// (DSN) for PostgreSQL connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		DBHost, DBUser, DBPassword, DBName, DBPort, DBSSLMode)

	// Database connection setup
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	// Establish database connection
	db, err := connectToDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// set up config
	var app = Config{
		DB: db,
	}

	ServicePort := os.Getenv("AUTHENTICATION_SERVICE_PORT")
	ServiceName := os.Getenv("AUTHENTICATION_SERVICE_NAME")

	if ServicePort == "" || ServiceName == "" {
		log.Fatal("Error: Service environment variables are not set")
	}

	log.Printf("%s is running on port: %s", ServiceName, ServicePort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", ServicePort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic()
	}
}
