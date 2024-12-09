package main

import (
	"log"

	"github.com/streadway/amqp"
)

type Message struct {
	Type string `json:"type"` 
	Body string `json:"body"` 
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5673/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare the exchange
	err = ch.ExchangeDeclare(
		"direct_logs", 
		"direct",      
		true,          
		false,         // Auto-deleted
		false,         // Internal
		false,         // No-wait
		nil,           // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}

	messages := []Message{
		{"type1", "This is a type1 message"},
		{"type2", "This is a type2 message"},
		{"type1", "Another type1 message"},
		{"type2", "Another type2 message"},
	}

	for _, msg := range messages {
		err = ch.Publish(
			"direct_logs",   // Exchange name
			msg.Type,        // Routing key
			false,           // Mandatory
			false,           // Immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg.Body),
			},
		)
		if err != nil {
			log.Printf("Failed to publish message: %v", err)
		} else {
			log.Printf("[x] Sent %s: %s", msg.Type, msg.Body)
			
		}
	}
}
