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
	// Database connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	// Connect to PostgreSQL (using "postgres" as dbname to connect to default db)
	var err error
	DB, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{TranslateError: true},
	)

	if err != nil {
		log.Fatalf("failed to establish database connection: %v", err)
	}

	// Check if the "bounting" database exists
	if err := createDatabaseIfNotExists(); err != nil {
		log.Fatalf("failed to create or connect to database: %v", err)
	}

	// Switch to the "bounting" database for further operations
	dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	DB, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{TranslateError: true},
	)

	if err != nil {
		log.Fatalf("failed to switch to the 'bounting' database: %v", err)
	}
}

// createDatabaseIfNotExists checks if the database exists and creates it if not
func createDatabaseIfNotExists() error {
	// Connect using "postgres" database to check if the "bounting" DB exists
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	))
	if err != nil {
		return fmt.Errorf("unable to connect to the database: %v", err)
	}
	defer conn.Close(context.Background())

	// Check if "bounting" database exists
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
