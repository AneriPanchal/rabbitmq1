// consumer1.go
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
	

	msgs, err := ch.Consume(
		"que101", // Queue name
		"",       // Consumer tag
		false,    // Auto-acknowledge (false to allow manual ack)
		false,    // Exclusive
		false,    // No-local
		false,    // No-wait
		nil,      // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	fmt.Println(" [*] Consumer1 waiting for messages of type1. To exit press CTRL+C")
	for msg := range msgs {
		if msg.RoutingKey == "type1" { // Process only type1 messages
			fmt.Printf(" [x] Consumer1 received: %s\n", msg.Body)
			msg.Ack(false) // Acknowledge the message
		} else {
			msg.Nack(false, true) // Requeue other messages for other consumers
		}
	}
}
