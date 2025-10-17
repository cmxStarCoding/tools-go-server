package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

// WebSocket 升级器
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 保存用户与连接的映射
var (
	Clients   = make(map[uint]*websocket.Conn) // userID -> conn
	ClientsMu sync.RWMutex                     // 读写锁，防止并发问题
)

// WsHandler 处理 WebSocket 连接
func WsHandler(c *gin.Context) {
	userID := c.Value("UserId").(uint)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("升级 WebSocket 出错:", err)
		return
	}

	// 保存连接
	ClientsMu.Lock()
	Clients[userID] = conn
	ClientsMu.Unlock()

	fmt.Printf("✅ 用户 %s 已连接\n", userID)

	defer func() {
		ClientsMu.Lock()
		delete(Clients, userID)
		ClientsMu.Unlock()
		conn.Close()
		fmt.Printf("❌ 用户 %s 已断开\n", userID)
	}()

	// 循环接收消息
	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("读取错误:", err)
			break
		}
		fmt.Printf("收到用户 %s 的消息: %s\n", userID, msg)
		// 回复客户端
		conn.WriteMessage(mt, []byte("服务器已收到: "+string(msg)))
	}
}
