package handler

import (
	"be-api-go/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Clients struct {
	PhoneNumber string
	Conn        *websocket.Conn
	Pool        *Connections
}

type Connections struct {
	Clients    map[*Clients]bool
	Register   chan *Clients
	Unregister chan *Clients
	Broadcast  chan Message
}

type Message struct {
	Message string `json:"message"`
}

var poolConnections = make([]*Clients, 0)

type AI struct {
	GenerativeModel *genai.GenerativeModel
	Context         context.Context
}

type WebSocketConnection struct {
	*websocket.Conn
	PhoneNumber string
}

type SocketAiResponse struct {
	Result  *genai.GenerateContentResponse `json:"result"`
	Message string                         `json:"message"`
	Part    genai.Part                     `json:"part"`
}

type SocketChatWithAiPayload struct {
	*model.Chats
	Action     string `json:"action"`
	Prompt     string `json:"prompt"`
	PromptType string `json:"prompt_type"`
	Parameter  string `json:"parameter"`
}

type SocketPayload struct {
	*model.Chats
	Prompt string
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}

type SocketPayloadNotification struct {
	*model.Notifications
	Action      string `json:"action"`
	PhoneNumber string `json:"phone_number"`
	SellerId    int    `json:"seller_id"`
}

func NewConnectionPool() *Connections {
	return &Connections{
		Clients:    make(map[*Clients]bool),
		Register:   make(chan *Clients),
		Unregister: make(chan *Clients),
		Broadcast:  make(chan Message),
	}
}

func (c *Clients) Read() {
	defer func() {
		fmt.Println("closing connection from: ", c.PhoneNumber)
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var message Message
		err := c.Conn.ReadJSON(&message)
		if err != nil {
			log.Printf("error on reading json message %v", err)
			return

		}
		c.Pool.Broadcast <- message
		fmt.Println("Message Received: ", message)
	}
}
