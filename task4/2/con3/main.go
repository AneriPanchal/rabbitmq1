// consumer3.go
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

	// Declare a fanout exchange (same as in the producer)
	err = ch.ExchangeDeclare(
		"fanout_exchange", // Name
		"fanout",          // Type
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare the exchange: %s", err)
	}

	// Declare an anonymous queue
	q, err := ch.QueueDeclare(
		"",    // Name (empty string for an auto-generated queue)
		false, // Durable
		false, // Delete when unused
		true,  // Exclusive to this connection
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// Bind the anonymous queue to the fanout exchange
	err = ch.QueueBind(
		q.Name,
		"",                // Routing key (ignored for fanout)
		"fanout_exchange", // Exchange name
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind the queue: %s", err)
	}

	// Start consuming messages
	msgs, err := ch.Consume(
		q.Name, // Queue name
		"",     // Consumer tag
		true,   // Auto-acknowledge
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	for msg := range msgs {
		fmt.Printf(" [x] Received: %s\n", msg.Body)
	}
}
