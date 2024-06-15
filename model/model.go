package model

import (
	"gorm.io/gorm"
)

type Users struct {
	*gorm.Model
	Username string `json:"name"`
	Password string `json:"password"`
}

type Customers struct {
	*gorm.Model
	PhoneNumber string `json:"phone_number"`
	Name        string
}

type Sellers struct {
	*gorm.Model
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Category    string      `json:"category"`
	Products    []Products  `json:"products" gorm:"foreignKey:SellerId;references:ID"`
	ChatRooms   []ChatRooms `json:"chat_rooms" gorm:"foreignKey:SellerId;references:ID"`
}

type Products struct {
	*gorm.Model
	SellerId    int    `json:"seller_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type ChatRooms struct {
	*gorm.Model
	SellerId int     `json:"seller_id"`
	Chats    []Chats `json:"chats" gorm:"foreignKey:ChatRoomId;references:ID"`
}

type Chats struct {
	*gorm.Model
	ChatRoomId  int    `json:"chat_room_id"`
	SellerId    *int   `json:"seller_id"`
	Sender      string `json:"sender"`
	Receiver    string `json:"receiver"`
	PhoneNumber string `json:"phone_number"`
	Chat        string `gorm:"type:text"`
}
