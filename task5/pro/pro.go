package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5673/") // Default port is 5672; adjust if necessary
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare a headers exchange
	err = ch.ExchangeDeclare(
		"header_exchange", // Exchange name
		"headers",         // Exchange type
		true,              // Durable
		false,             // Auto-deleted
		false,             // Internal
		false,             // No-wait
		nil,               // Arguments
	)
	failOnError(err, "Failed to declare exchange")

	// Publish "type 1" message
	headers1 := amqp.Table{"type": "1"} // Headers for type 1
	body1 := "Message of type 1"
	err = ch.Publish(
		"header_exchange", // Exchange name
		"",                // Routing key (ignored for headers exchange)
		false,             // Mandatory
		false,             // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body1),
			Headers:     headers1, // Add headers
		},
	)
	failOnError(err, "Failed to publish type 1 message")
	log.Printf(" [x] Sent: %s", body1)

	// Publish "type 2" message
	headers2 := amqp.Table{"type": "2"} // Headers for type 2
	body2 := "Message of type 2"
	err = ch.Publish(
		"header_exchange", // Exchange name
		"",                // Routing key (ignored for headers exchange)
		false,             // Mandatory
		false,             // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body2),
			Headers:     headers2, // Add headers
		},
	)
	failOnError(err, "Failed to publish type 2 message")
	log.Printf(" [x] Sent: %s", body2)
}