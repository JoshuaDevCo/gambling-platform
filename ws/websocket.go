// websocket/server.go
package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WebSocketServer struct {
	mu       sync.Mutex
	clients  map[*websocket.Conn]bool
	upgrader websocket.Upgrader
}

// NewWebSocketServer creates a new instance of WebSocketServer.
func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		clients:  make(map[*websocket.Conn]bool),
		upgrader: websocket.Upgrader{},
	}
}

// HandleConnections handles WebSocket connections.
func (ws *WebSocketServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Register the new client
	ws.mu.Lock()
	ws.clients[conn] = true
	ws.mu.Unlock()

	// Listen for messages from the client
	for {
		// Read the message from the client
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			ws.mu.Lock()
			delete(ws.clients, conn)
			ws.mu.Unlock()
			break
		}

		// Handle different types of messages (subscribe, unsubscribe, etc.)
		switch messageType {
		case websocket.TextMessage:
			// Handle text messages (JSON-encoded commands)
			ws.handleTextMessage(p)
		case websocket.CloseMessage:
			// Handle close messages
			ws.mu.Lock()
			delete(ws.clients, conn)
			ws.mu.Unlock()
			break
		}
	}
}

// handleTextMessage handles incoming text messages from clients.
func (ws *WebSocketServer) handleTextMessage(message []byte) {
	// Handle different types of commands or messages
	// For example, you can implement logic for subscribing/unsubscribing to events.
	// You can define a protocol for messages (e.g., JSON format) that includes the type of message and necessary data.
	// Here, we're simply logging the received message.
	log.Printf("Received message: %s", string(message))
}

// Broadcast sends a message to all connected clients.
func (ws *WebSocketServer) Broadcast(messageType int, message []byte) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	// Send the message to all connected clients
	for conn := range ws.clients {
		err := conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println(err)
			conn.Close()
			delete(ws.clients, conn)
		}
	}
}
