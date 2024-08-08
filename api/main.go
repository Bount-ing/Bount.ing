package main

import (
	"log"
	"open-bounties-api/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v74"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	if stripe.Key == "" {
		log.Fatalf("Error loading .env file: STRIPE_SECRET_KEY not found")
	}
	r := routes.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
