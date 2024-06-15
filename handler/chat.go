package handler

import (
	"be-api-go/database"
	"be-api-go/model"
	"fmt"
	//"log"
)

//func handleChatWithUser(w http.ResponseWriter, r *http.Request) {
//	conn, err := upgrader.Upgrade(w, r, nil)
//
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	defer conn.Close()
//
//	phoneNumber := r.URL.Query().Get("phoneNumber")
//
//	//currentConn := WebSocketConnection{Conn: conn, PhoneNumber: phoneNumber}
//	//connections = append(connections, &currentConn)
//
//	//handleChat(&currentConn, connections)
//	return nil
//}

func InsertAndSaveChat(payload *model.Chats) error {
	//find chat room
	var chatRoom model.ChatRooms
	if err := database.DB.Where("id = ?", payload.ChatRoomId).First(&chatRoom).Error; err != nil {
		fmt.Println("Error on finding chat room: ", err)
		return err
	}

	if err := database.DB.Create(&payload).Error; err != nil {
		fmt.Println("Error on inserting chat: ", err)
		return err
	}

	return nil
}
