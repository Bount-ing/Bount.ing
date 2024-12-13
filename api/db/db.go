package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// Explicit extraction of environment variables
	dbName := strings.Trim(os.Getenv("POSTGRES_DB"), "\"")
	dbUser := strings.Trim(os.Getenv("POSTGRES_USER"), "\"")
	dbHost := strings.Trim(os.Getenv("POSTGRES_HOST"), "\"")
	dbPassword := strings.Trim(os.Getenv("POSTGRES_PASSWORD"), "\"")
	dbPort := strings.Trim(os.Getenv("POSTGRES_PORT"), "\"")

	// First connection string (to default postgres database)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbPort,
	)

	// Connect to PostgreSQL
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

	// Check if the database exists and create if not
	if err := createDatabaseIfNotExists(dbName, dsn); err != nil {
		log.Fatalf("failed to create or connect to database: %v", err)
	}

	// Connection string for the specific database
	dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	// Reconnect to the specific database
	DB, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{TranslateError: true},
	)
	if err != nil {
		log.Fatalf("failed to switch to the database: %v", err)
	}
}

func createDatabaseIfNotExists(dbName string, dsn string) error {
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return fmt.Errorf("unable to connect to the database: %v", err)
	}
	defer conn.Close(context.Background())

	// Check if database exists
	var exists bool
	err = conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", dbName).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if database exists: %v", err)
	}

	// If it doesn't exist, create it
	if !exists {
		// Use a parameterized query to safely create the database
		_, err := conn.Exec(context.Background(), fmt.Sprintf("CREATE DATABASE %s", pq.QuoteIdentifier(dbName)))
		if err != nil {
			return fmt.Errorf("failed to create database: %v", err)
		}
	}
	return nil
}
