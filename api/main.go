package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/bount-ing/bount.ing/api/processor"
	"github.com/bount-ing/bount.ing/api/routes"
)

func checkEnv() {
	// Mandatory environment variables
	// General
	if os.Getenv("APP_BASE_URL") == "" {
		log.Fatal("APP_BASE_URL is not set")
	}
	if os.Getenv("API_BASE_URL") == "" {
		log.Fatal("API_BASE_URL is not set")
	}
	if os.Getenv("JWT_SECRET_KEY") == "" {
		log.Fatal("JWT_SECRET_KEY is not set")
	}

	// Database
	if os.Getenv("DB_HOST") == "" {
		log.Fatal("DB_HOST is not set")
	}
	if os.Getenv("DB_PORT") == "" {
		log.Fatal("DB_PORT is not set")
	}
	if os.Getenv("DB_USER") == "" {
		log.Fatal("DB_USER is not set")
	}
	if os.Getenv("DB_PWD") == "" {
		log.Fatal("DB_PWD is not set")
	}
	if os.Getenv("DB_NAME") == "" {
		log.Fatal("DB_NAME is not set")
	}

	// Github
	if os.Getenv("GITHUB_CLIENT_SECRET") == "" {
		log.Fatal("GITHUB_CLIENT_SECRET is not set")
	}
	if os.Getenv("GITHUB_CLIENT_ID") == "" {
		log.Fatal("GITHUB_CLIENT_ID is not set")
	}
	if os.Getenv("GITHUB_REDIRECT_URL") == "" {
		log.Fatal("GITHUB_REDIRECT_URL is not set")
	}
	if os.Getenv("GITHUB_REDIRECT_URI") == "" {
		log.Fatal("GITHUB_REDIRECT_URI is not set")
	}
	if os.Getenv("GITHUB_WEBHOOK_SECRET") == "" {
		log.Fatal("GITHUB_WEBHOOK_SECRET is not set")
	}

	// GMail
	if os.Getenv("NOREPLY_MAIL_ADDRESS") == "" {
		log.Fatal("NOREPLY_MAIL_ADDRESS is not set")
	}
	if os.Getenv("NOREPLY_MAIL_PASSWD") == "" {
		log.Fatal("NOREPLY_MAIL_PASSWD is not set")
	}
	if os.Getenv("CONTACT_MAIL_ADDRESS") == "" {
		log.Fatal("CONTACT_MAIL_ADDRESS is not set")
	}
	if os.Getenv("SUPPORT_MAIL_ADDRESS") == "" {
		log.Fatal("SUPPORT_MAIL_ADDRESS is not set")
	}

	// Stripe
	if os.Getenv("STRIPE_SECRET_KEY") == "" {
		log.Fatal("STRIPE_SECRET_KEY is not set")
	}
	if os.Getenv("STRIPE_CLIENT_ID") == "" {
		log.Fatal("STRIPE_CLIENT_ID is not set")
	}
	if os.Getenv("STRIPE_REDIRECT_URI") == "" {
		log.Fatal("STRIPE_REDIRECT_URI is not set")
	}
	if os.Getenv("STRIPE_REDIRECT_URL") == "" {
		log.Fatal("STRIPE_REDIRECT_URL is not set")
	}

	// RabbitMQ
	if os.Getenv("RABBITMQ_HOST") == "" {
		log.Fatal("RABBITMQ_HOST is not set")
	}
	if os.Getenv("RABBITMQ_PORT") == "" {
		log.Fatal("RABBITMQ_PORT is not set")
	}
	if os.Getenv("RABBITMQ_USER") == "" {
		log.Fatal("RABBITMQ_USER is not set")
	}
	if os.Getenv("RABBITMQ_PASSWD") == "" {
		log.Fatal("RABBITMQ_PASSWD is not set")
	}

	// Optional environment variables
	// Discord
	if os.Getenv("DISCORD_BOUNTIES_WEBHOOK_URL") == "" {
		log.Println("DISCORD_BOUNTIES_WEBHOOK_URL is not set")
	}

	if os.Getenv("DISCORD_ERRORS_WEBHOOK_URL") == "" {
		log.Println("DISCORD_ERRORS_WEBHOOK_URL is not set")
	}

	if os.Getenv("DISCORD_EVENTS_WEBHOOK_URL") == "" {
		log.Println("DISCORD_EVENTS_WEBHOOK_URL is not set")
	}
}

func main() {
	// Check environment variables
	checkEnv()

	// Create a WaitGroup to manage our goroutines
	var wg sync.WaitGroup

	// Create a context that can be canceled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the RabbitMQ consumer in a background goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Starting RabbitMQ consumer...")
		err := processor.ProcessEvents()
		if err != nil {
			log.Printf("Consumer error: %v", err)
		}
	}()

	// Set up signal handling for graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Start the web server in a separate goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		r := routes.SetupRouter()
		// Use a custom server so we can shut it down gracefully
		server := &http.Server{
			Addr:    ":8080",
			Handler: r,
		}

		// Listen for the shutdown signal in a separate goroutine
		go func() {
			<-ctx.Done()
			log.Println("Shutting down HTTP server...")
			// Allow 5 seconds for server to finish processing requests
			shutdownCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			if err := server.Shutdown(shutdownCtx); err != nil {
				log.Printf("HTTP server shutdown error: %v", err)
			}
		}()

		log.Println("Starting HTTP server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server error: %v", err)
		}
		log.Println("HTTP server stopped")
	}()

	// Wait for termination signal
	<-shutdown
	log.Println("Shutdown signal received")

	// Cancel context to signal all goroutines to shut down
	cancel()

	// Wait for all goroutines to finish
	wg.Wait()
	log.Println("All services shut down, exiting")
}
