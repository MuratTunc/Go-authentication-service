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
	DB *gorm.DB
}

var (
	DBHost     = os.Getenv("DB_HOST")     // "auth-db" from docker-compose
	DBUser     = os.Getenv("DB_USER")     // "auth_user"
	DBPassword = os.Getenv("DB_PASSWORD") // "auth_password"
	DBName     = os.Getenv("DB_NAME")     // "auth_db"
	DBPort     = os.Getenv("DB_PORT")     // "5432"
)

// connectToDB function to connect to PostgreSQL using constants
func connectToDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DBHost, DBUser, DBPassword, DBName, DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate to create tables
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db, nil
}

func main() {
	PrintEnvVariables()

	db, err := connectToDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	var app = Config{DB: db}

	ServicePort := "8080"
	ServiceName := "AUTHENTICATION SERVICE"

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
