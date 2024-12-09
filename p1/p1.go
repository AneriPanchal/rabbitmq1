// producer.go
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

	//  exchange
	err = ch.ExchangeDeclare(
		"direct_exchange", // Name
		"direct",          // Type
		true,              // Durable
		false,             // Auto-deleted
		false,             // Internal
		false,             // No-wait
		nil,               // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %s", err)
	}

	// Declare  queue
	_, err = ch.QueueDeclare(
		"que101", 
		true,     
		false,    
		false,   
		false,    
		nil,      
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// Bind the queue to the exchange 
	err = ch.QueueBind(
		"que101",           // Queue name
		"type1",            // Routing key
		"direct_exchange",  // Exchange name
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind the queue with type1: %s", err)
	}

	err = ch.QueueBind(
		"que101",           // Queue name
		"type2",            // Routing key
		"direct_exchange",  // Exchange name
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind the queue with type2: %s", err)
	}

	// Publish messages
	err = ch.Publish(
		"direct_exchange", // Exchange
		"type1",           // Routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Message for Type1"),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish type1 message: %s", err)
	}
	fmt.Println(" [x] Sent 'Message for Type1'")

	err = ch.Publish(
		"direct_exchange", // Exchange
		"type2",           // Routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Message for Type2"),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish type2 message: %s", err)
	}
	fmt.Println(" [x] Sent 'Message for Type2'")
}
