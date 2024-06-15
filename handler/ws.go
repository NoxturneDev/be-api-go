package handler

import (
	"be-api-go/model"
	"context"
	"encoding/json"
	"github.com/google/generative-ai-go/genai"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
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

//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"google.golang.org/api/iterator"
//	"google.golang.org/api/option"
//	"log"
//	"os"
//
//	"github.com/google/generative-ai-go/genai"
//	"github.com/gorilla/websocket"
//	"github.com/joho/godotenv"
//)
//
//func WebsocketHandler(c *websocket.Conn) error {
//	// c.Locals is added to the *websocket.Conn
//	log.Println(c.Locals("allowed"))  // true
//	log.Println(c.Params("id"))       // 123
//	log.Println(c.Query("v"))         // 1.0
//	log.Println(c.Cookies("session")) // ""
//
//	// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
//	var (
//		mt  int
//		msg []byte
//		err error
//	)
//	for {
//		if mt, msg, err = c.ReadMessage(); err != nil {
//			log.Println("read:", err)
//			break
//		}
//		log.Printf("recv: %s", msg)
//
//		if err = c.WriteMessage(mt, msg); err != nil {
//			log.Println("write:", err)
//			break
//		}
//	}
//
//	return nil
//}
//
////func AiWebsocketHandler(c *websocket.Conn) error {
////	log.Println(c.Locals("allowed"))  // true
////	log.Println(c.Params("id"))       // 123
////	log.Println(c.Query("v"))         // 1.0
////	log.Println(c.Cookies("session")) // ""
////
////	// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
////	var (
////		mt  int
////		msg []byte
////		err error
////	)
////
////	for {
////		if mt, msg, err = c.ReadMessage(); err != nil {
////			log.Println("read:", err)
////			break
////		}
////		log.Printf("recv: %s", msg)
////
////		if err = c.WriteMessage(mt, msg); err != nil {
////			log.Println("write:", err)
////			break
////		}
////	}
////
////	err = godotenv.Load()
////	if err != nil {
////		log.Fatal("Error loading .env file")
////	}
////
////	geminiApiKey := os.Getenv("GEMINI_API_KEY")
////
////	// Access your API key as an environment variable (see "Set up your API key" above)
////	mdl.Context = context.Background()
////	client, err := genai.NewClient(mdl.Context, option.WithAPIKey(geminiApiKey))
////	if err != nil {
////		fmt.Println(err)
////	}
////
////	defer client.Close()
////	mdl.GenerativeModel = client.GenerativeModel("gemini-1.5-flash")
////
////	iter := mdl.GenerativeModel.GenerateContentStream(mdl.Context, genai.Text("Tell me a story about a lumberjack and his giant ox. Keep it very short."))
////	for {
////		_, err := iter.Next()
////		if errors.Is(err, iterator.Done) {
////			break
////		}
////		if err != nil {
////			return err
////		}
////
////	}
////
////	return nil
////}
//
//func AiWebsocketHandler(c *websocket.Conn) error {
//	log.Println(c.Locals("allowed"))  // true
//	log.Println(c.Params("id"))       // 123
//	log.Println(c.Query("v"))         // 1.0
//	log.Println(c.Cookies("session")) // ""
//
//	// Load environment variables
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//
//	// Set up API key and client
//	geminiApiKey := os.Getenv("GEMINI_API_KEY")
//	ctx := context.Background()
//	client, err := genai.NewClient(ctx, option.WithAPIKey(geminiApiKey))
//	if err != nil {
//		fmt.Println(err)
//		return c.WriteMessage(websocket.TextMessage, []byte("Failed to create AI client"))
//	}
//	defer client.Close()
//
//	// Set up the model
//	//model := client.GenerativeModel("gemini-1.5-flash")
//
//	// Generate content stream
//
//	c.Locals()
//	//iter := model.GenerateContentStream(ctx, genai.Text("Tell me a story about a lumberjack and his giant ox. Keep it very short."))
//	//for {
//	//	resp, err := iter.Next()
//	//	if errors.Is(err, iterator.Done) {
//	//		break
//	//	}
//	//	if err != nil {
//	//		log.Println("stream error:", err)
//	//		return c.WriteMessage(websocket.TextMessage, []byte("Error occurred during stream"))
//	//	}
//	//
//	//	// Send each response chunk to the WebSocket client
//	//	log.Println("Sending message:", resp)
//	//	//if err := c.WriteMessage(websocket.TextMessage, message); err != nil {
//	//	//	log.Println("write error:", err)
//	//	//	break
//	//	//}
//	//}
//
//	return nil
//}
