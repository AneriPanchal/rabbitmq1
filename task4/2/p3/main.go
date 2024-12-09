// producer1.go
package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5673/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	// Declare a fanout exchange
	err = ch.ExchangeDeclare(
		"fanout_exchange", // Name
		"fanout",          // Type
		true,              // Durable
		false,             // Auto-deleted
		false,             // Internal
		false,             // No-wait
		nil,               // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare the exchange: %s", err)
	}

	// Publish a message
	body := "Message from Producer 3"
	err = ch.Publish(
		"fanout_exchange", // Exchange name
		"",                // Routing key (ignored for fanout)
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}

	fmt.Println(" [x] Sent:", body)
}
