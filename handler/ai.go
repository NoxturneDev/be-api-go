package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var mdl AI

func WsAiHandler(connections *Connections, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	phoneNumber := r.URL.Query().Get("phoneNumber")
	client := &Clients{
		Conn:        conn,
		PhoneNumber: phoneNumber,
		Pool:        connections,
	}

	fmt.Println("Coming connection from ", phoneNumber)
	connections.Register <- client
	client.Read()
}

func AiChatHandler(connections *Connections) {
	for {
		select {
		case client := <-connections.Register:

			// check if there's existing phone number inside the pool connection
			// if there is, then remove it
			for c := range connections.Clients {
				log.Println("incoming phone number: ", client.PhoneNumber)
				log.Println("existing phone number", c.PhoneNumber)
				if client.PhoneNumber == c.PhoneNumber {
					log.Println("there's existing phone number in the pool connection")
					break
				}
			}

			connections.Clients[client] = true
			poolConnections = append(poolConnections, client)
			log.Println("Size of Connection Pool: ", len(poolConnections))

			for client := range connections.Clients {
				client.Conn.WriteJSON(Message{
					Message: "New User Connected" + client.PhoneNumber,
				})
			}
			break
		case client := <-connections.Unregister:
			delete(connections.Clients, client)
			removeClient(client)
			log.Println("Size of Connection Pool After Closing: ", len(poolConnections))
			for client := range connections.Clients {
				client.Conn.WriteJSON(Message{
					Message: "User Disconnected" + client.PhoneNumber,
				})
			}
			break
		case message := <-connections.Broadcast:
			log.Println("Sending message to all clients in the pool")
			for _, client := range poolConnections {
				if err := client.Conn.WriteJSON(message); err != nil {
					log.Println("something went wrong while sending message to the client", err)
					if strings.Contains(err.Error(), "websocket: close sent") {
						connections.Unregister <- client
					}
					return
				}
			}
		}
	}
}

func removeClient(c *Clients) {
	index := -1
	for i, client := range poolConnections {
		if client.PhoneNumber == c.PhoneNumber {
			index = i
			break
		}
	}

	if index != -1 {
		poolConnections = append(poolConnections[:index], poolConnections[index+1:]...)
	}
}

//func handleChat(client *Clie nts) {
//	client.Conn.WriteJSON(Message{
//		Message: message,
//	})
//}

//func sendAiResultInStreamParts(currentConn *WebSocketConnection, part
//genai.Part) {
//currentConn.WriteJSON(SocketAiResponse{
//Part: part,
//})
//}

//func handleChat(client *Clients) {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//
//	defer func() {
//		if r := recover(); r != nil {
//			log.Println("ERROR: ", fmt.Sprintf("%v", r))
//		}
//	}()
//
//	//Set up API key and client
//	geminiApiKey := os.Getenv("GEMINI_API_KEY")
//	ctx := context.Background()
//	client, err := genai.NewClient(ctx, option.WithAPIKey(geminiApiKey))
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	// TODO: create chat session
//	model := client.GenerativeModel("gemini-1.5-flash")
//	cs := model.StartChat()
//
//	for {
//		payload := SocketChatWithAiPayload{}
//		err := currentConn.ReadJSON(&payload)
//		if err != nil {
//			if strings.Contains(err.Error(), "websocket: close") {
//				sendStringMessage(currentConn, "LEAVE")
//				return
//			}
//
//			log.Println("ERROR", err.Error())
//			continue
//		}
//
//		//InsertAndSaveChat(payload.Chats)
//		if payload.PromptType == "suggestion" {
//			payload.Prompt = util.CreateSuggestionAIPrompt(payload.Prompt)
//		} else if payload.PromptType == "translate" {
//			payload.Prompt = util.TranslateToLanguagePrompt(payload.Prompt)
//		} else if payload.PromptType == "summary" {
//			payload.Prompt = util.CreateSummaryAiPrompt(payload.Prompt)
//		} else if payload.PromptType == "sentimen" {
//			payload.Prompt = util.CreateCustomerSentimentAnalysisPrompt(payload.Prompt)
//		} else if payload.PromptType == "user" {
//			payload.Prompt = util.CreateMockUserChatting(payload.Prompt)
//		}
//
//		fmt.Println(payload.Prompt)
//
//		iter := cs.SendMessageStream(ctx, genai.Text(payload.Prompt))
//		for {
//			resp, err := iter.Next()
//			if errors.Is(err, iterator.Done) {
//				//TODO:  save this in the database when all the chat is done
//				fmt.Println(iter.MergedResponse().Candidates[0].Content.Parts[0])
//				currentConn.WriteJSON(iter.MergedResponse().Candidates[0].Content.Parts[0])
//				break
//			}
//			if err != nil {
//				log.Println("stream error:", err)
//				sendStringMessage(currentConn, "Failed to generate content")
//			}
//
//			// Send each response chunk to the WebSocket client
//			log.Println("Sending message:", resp)
//			var temp = iter.MergedResponse().Candidates[0].Content.Parts[0]
//			sendAiResultInStreamParts(currentConn, temp)
//		}
//	}
//
//}
