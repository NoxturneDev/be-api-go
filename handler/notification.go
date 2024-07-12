package handler

//func NotificationHandler(w http.ResponseWriter, r *http.Request) {
//	conn, err := upgrader.Upgrade(w, r, nil)
//
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	phoneNumber := r.URL.Query().Get("phoneNumber")
//	currentConn := WebSocketConnection{Conn: conn, PhoneNumber: phoneNumber}
//	connections = append(connections, &currentConn)
//
//	fmt.Println("New connection from", phoneNumber)
//	payload := SocketPayloadNotification{}
//	err = currentConn.ReadJSON(&payload)
//	if err != nil {
//		log.Println("ERROR read payload: ", err.Error())
//		return
//	}
//
//	fmt.Println("Payload: ", payload)
//	switch payload.Action {
//	case "create":
//		fmt.Println("create notification....")
//		CreateNotification(&currentConn, payload)
//	case "retrieve":
//		//handleRetrieve(conn, msg.Data)
//	default:
//		log.Println("unknown action:", payload.Action)
//		conn.WriteJSON(map[string]string{"error": "unknown action"})
//	}
//	return
//}
//
//func CreateNotification(currentConn *WebSocketConnection, payload SocketPayloadNotification) {
//	// insert notification to database
//	newNotification := model.Notifications{
//		PhoneNumber: payload.PhoneNumber,
//		SellerId:    payload.SellerId,
//	}
//
//	if err := database.DB.Create(&newNotification).Error; err != nil {
//		log.Println("ERROR insert to db: ", err.Error())
//		return
//	}
//
//	err := currentConn.WriteJSON(map[string]string{"message": "notification created"})
//	if err != nil {
//		log.Println("ERROR sending data: ", err.Error())
//		return
//	}
//}
//
//func GetNotificationByPhoneNumber(phoneNumber string) ([]model.Notifications, error) {
//	var notifications []model.Notifications
//	if err := database.DB.Where("phone_number = ?", phoneNumber).Find(&notifications).Error; err != nil {
//		return nil, err
//	}
//
//	return notifications, nil
//}
