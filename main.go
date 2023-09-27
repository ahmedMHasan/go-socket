package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	// Initialize a Go module
	// go mod init go-socket

	// Set up a graceful shutdown channel
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Create a RabbitMQ connection
	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open RabbitMQ channel: %v", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"notification_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to consume from RabbitMQ: %v", err)
	}

	router := gin.Default()

	router.GET("/ws", func(c *gin.Context) {
		// Upgrade HTTP request to WebSocket connection
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("WebSocket upgrade failed: %v", err)
			return
		}
		defer ws.Close()

		for msg := range msgs {
			// Write RabbitMQ messages to the WebSocket
			err := ws.WriteMessage(websocket.TextMessage, msg.Body)
			if err != nil {
				log.Printf("WebSocket write error: %v", err)
				return
			}
		}
	})

	// Start the WebSocket server in a goroutine
	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to start WebSocket server: %v", err)
		}
	}()

	// Wait for a shutdown signal
	<-shutdown
	fmt.Println("Shutting down gracefully...")

	// You can add cleanup logic here

	fmt.Println("Server gracefully stopped.")
}
