package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

func AddConnection(userID uint, conn *websocket.Conn) {
	ClientsMu.Lock()
	defer ClientsMu.Unlock()

	if c := UserIdToConn[userID]; c != nil {
		c.Close()
	}
	UserIdToConn[userID] = conn
	ConnToUserId[conn] = userID
}

func KeepAlive(conn *websocket.Conn) {
	Timer := time.NewTicker(HeartbeatInterval * time.Second)
	UserId := ConnToUserId[conn]

	defer Timer.Stop()

	for {
		select {
		case <-Timer.C:
			//判断心跳时间
			value, ok := ConnHeart[UserId]
			if ok {
				if time.Now().Unix()-value > HeartbeatInterval {
					fmt.Printf("用户id值 %d心跳超时，已断开链接\n", UserId)
					CloseConn(conn)
					return
				}
			} else {
				CloseConn(conn)
			}
		}
	}
}
