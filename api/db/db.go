package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4" // Use pgx to directly check and create database
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// Database connection string for the initial setup
	initialDsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		"postgres", // Use default database for initial setup
		os.Getenv("POSTGRES_PORT"),
	)

	// Connect to PostgreSQL (using "postgres" as dbname to connect to default db)
	var err error
	DB, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  initialDsn,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{TranslateError: true},
	)

	if err != nil {
		log.Fatalf("failed to establish database connection: %v", err)
	}

	// Check if the target database exists and create it if necessary
	if err := createDatabaseIfNotExists(); err != nil {
		log.Fatalf("failed to create or connect to database: %v", err)
	}

	// Switch to the target database for further operations
	targetDsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	DB, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  targetDsn,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{TranslateError: true},
	)

	if err != nil {
		log.Fatalf("failed to switch to the target database: %v", err)
	}
}

// createDatabaseIfNotExists checks if the database exists and creates it if not
func createDatabaseIfNotExists() error {
	// Use the default database for administrative tasks
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		"postgres", // Use default database
		os.Getenv("POSTGRES_PORT"),
	))
	if err != nil {
		return fmt.Errorf("unable to connect to the default database: %v", err)
	}
	defer conn.Close(context.Background())

	// Check if the target database exists
	var exists bool
	err = conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", os.Getenv("POSTGRES_DB")).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if database exists: %v", err)
	}

	// If it doesn't exist, create it
	if !exists {
		_, err := conn.Exec(context.Background(), fmt.Sprintf("CREATE DATABASE %s", os.Getenv("POSTGRES_DB")))
		if err != nil {
			return fmt.Errorf("failed to create database: %v", err)
		}
	}

	return nil
}
