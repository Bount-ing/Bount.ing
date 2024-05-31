package main

import (
	"log"
	"open-bounties-api/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	r := routes.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
