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
		"que101",
		"",
		false, // Auto-acknowledge (false to allow manual ack)
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	fmt.Println(" [*] Consumer2 waiting for messages of type2. To exit press CTRL+C")
	for msg := range msgs {
		if msg.RoutingKey == "type2" {
			fmt.Printf(" [x] Consumer2 received: %s\n", msg.Body)
			msg.Ack(false) // Acknowledge the message
		} else {
			msg.Nack(false, true)
		}
	}
}
