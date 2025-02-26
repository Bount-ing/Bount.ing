package processor

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func ProcessEvents() error {
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	rabbitmqPort := os.Getenv("RABBITMQ_PORT")
	rabbitmqUser := os.Getenv("RABBITMQ_USER")
	rabbitmqPassword := os.Getenv("RABBITMQ_PASSWD")

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitmqUser, rabbitmqPassword, rabbitmqHost, rabbitmqPort))
	if err != nil {
		log.Println("failed to connect to RabbitMQ:", err)
		time.Sleep(30 * time.Second)
		return ProcessEvents()
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}
	defer ch.Close()

	// Declare a queue
	q, err := ch.QueueDeclare(
		"events", // queue name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}

	// Create a consumer
	msgs, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer tag
		true,   // auto-acknowledge
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	// Create a channel to handle application shutdown
	forever := make(chan bool)

	// Process messages in a goroutine
	go func() {
		for msg := range msgs {
			fmt.Printf("Received a message: %s\n", msg.Body)
			if err := processor(msg.Body); err != nil {
				RepublishEvent("events", msg.Body)
			}
		}
	}()

	log.Println("Consumer started. Waiting for messages...")

	// Block indefinitely
	<-forever

	return nil
}

func processor(msg []byte) error {
	fmt.Println("Processing message:", string(msg))

	//Unmarshal the message
	message := map[string]interface{}{}
	err := json.Unmarshal(msg, &message)
	if err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	if message["event"] == "user_creation" {
		if err := processUserCreation(message); err != nil {
			return err
		}
	}
	return nil
}
