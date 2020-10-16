package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"time"
)

//It keeps a list of clients those are currently attached
//and broadcasting events to those clients.
type Event struct {
	// Events are pushed to this channel by the main events-gathering routine
	Message chan string

	// New client connections
	NewClients chan chan string

	// Closed client connections
	ClosedClients chan chan string

	// Total client connections
	TotalClients map[chan string]bool
}

// New event messages are broadcast to all registered client connection channels
type ClientChan chan string

func main() {
	router := gin.Default()

	// Add event-streaming headers
	router.Use(HeadersMiddleware())

	// Initialize new streaming server
	stream := NewServer()
	router.Use(stream.serveHTTP())

	// Basic Authentication
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "admin123", // username : admin, password : admin123
	}))

	// Authorized client can stream the event
	authorized.GET("/stream", func(c *gin.Context) {
		// We are streaming current time to clients in the interval 10 seconds
		go func() {
			for {
				time.Sleep(time.Second * 10)
				now := time.Now().Format("2006-01-02 15:04:05")
				currentTime := fmt.Sprintf("The Current Time Is %v", now)

				// Send current time to clients message channel
				stream.Message <- currentTime
			}
		}()

		c.Stream(func(w io.Writer) bool {
			// Stream message to client from message channel
			if msg, ok := <-stream.Message; ok {
				c.SSEvent("message", msg)
				return true
			}
			return false
		})
	})

	//Parse Static files
	router.StaticFile("/", "./public/index.html")

	router.Run(":8085")
}

// Initialize event and Start procnteessing requests
func NewServer() (event *Event) {

	event = &Event{
		Message:       make(chan string),
		NewClients:    make(chan chan string),
		ClosedClients: make(chan chan string),
		TotalClients:  make(map[chan string]bool),
	}

	go event.listen()

	return
}

//It Listens all incoming requests from clients.
//Handles addition and removal of clients and broadcast messages to clients.
func (stream *Event) listen() {
	for {
		select {
		// Add new available client
		case client := <-stream.NewClients:
			stream.TotalClients[client] = true
			log.Printf("Client added. %d registered clients", len(stream.TotalClients))

		// Remove closed client
		case client := <-stream.ClosedClients:
			delete(stream.TotalClients, client)
			log.Printf("Removed client. %d registered clients", len(stream.TotalClients))

		// Broadcast message to client
		case eventMsg := <-stream.Message:
			for clientMessageChan := range stream.TotalClients {
				clientMessageChan <- eventMsg
			}
		}
	}
}

func (stream *Event) serveHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Initialize client channel
		clientChan := make(ClientChan)

		// Send new connection to event server
		stream.NewClients <- clientChan

		defer func() {
			// Send closed connection to event server
			stream.ClosedClients <- clientChan
		}()

		go func() {
			// Send connection that is closed by client to event server
			<-c.Done()
			stream.ClosedClients <- clientChan
		}()

		c.Next()
	}
}

func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Next()
	}
}
