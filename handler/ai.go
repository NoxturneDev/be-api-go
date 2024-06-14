package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
	"strings"
)

var mdl AI

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var connections = make([]*WebSocketConnection, 0)

type AI struct {
	GenerativeModel *genai.GenerativeModel
	Context         context.Context
}

type WebSocketConnection struct {
	*websocket.Conn
	PhoneNumber string
}

type SocketResponse struct {
	Message string
	Result  *genai.GenerateContentResponse
	Part    genai.Part
}

type SocketPayload struct {
	Prompt string
}

func AiChatHandler(w http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()

	phoneNumber := r.URL.Query().Get("phoneNumber")

	currentConn := WebSocketConnection{Conn: conn, PhoneNumber: phoneNumber}
	connections = append(connections, &currentConn)

	handleChat(&currentConn, connections)
	return nil
}

func handleChat(currentConn *WebSocketConnection, connections []*WebSocketConnection) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: ", fmt.Sprintf("%v", r))
		}
	}()

	// Set up API key and client
	geminiApiKey := os.Getenv("GEMINI_API_KEY")
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(geminiApiKey))
	if err != nil {
		fmt.Println(err)
	}

	// TODO: create chat session
	model := client.GenerativeModel("gemini-1.5-flash")
	cs := model.StartChat()

	for {
		payload := SocketPayload{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				sendStringMessage(currentConn, "LEAVE")
				return
			}

			log.Println("ERROR", err.Error())
			continue
		}

		iter := cs.SendMessageStream(ctx, genai.Text(payload.Prompt))
		for {
			resp, err := iter.Next()
			if errors.Is(err, iterator.Done) {
				//TODO:  save this in the database when all the chat is done
				fmt.Println(iter.MergedResponse().Candidates[0].Content.Parts[0])
				break
			}
			if err != nil {
				log.Println("stream error:", err)
				sendStringMessage(currentConn, "Failed to generate content")
			}

			// Send each response chunk to the WebSocket client
			log.Println("Sending message:", resp)
			sendAiResult(currentConn, resp.Candidates[0].Content.Parts[0])
		}
	}

}

func sendStringMessage(currentConn *WebSocketConnection, message string) {
	currentConn.WriteJSON(SocketResponse{
		Message: message,
	})
}

func sendAiResult(currentConn *WebSocketConnection, part genai.Part) {
	currentConn.WriteJSON(SocketResponse{
		Part: part,
	})
}
