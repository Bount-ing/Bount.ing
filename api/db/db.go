package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	var err error

	DB, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatalf("failed to establish database connection: %v", err)
	}
}
