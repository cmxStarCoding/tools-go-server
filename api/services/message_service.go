package services

import (
	"github.com/gorilla/websocket"
	"journey/ws"
	"log"
)

// MessageService 用户服务
type MessageService struct{}

func (s MessageService) SendMessage(message string) {

	log.Println("👀 Received11:", message)
	UserId := 17
	ws.ClientsMu.RLock()
	conn, ok := ws.Clients[uint(UserId)]
	ws.ClientsMu.RUnlock()
	if !ok {
		log.Printf("👀 用户id%d未连接websocket服务", UserId)
		return
	}
	SendMessage := `{"a":"b"}`
	// 发送消息
	err := conn.WriteMessage(websocket.TextMessage, []byte(SendMessage))
	if err != nil {
		log.Printf("👀 向用户id%d发送websocket消息%s成功", UserId, SendMessage)
	}
}
