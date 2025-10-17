package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

// WebSocket 升级器
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const IsStartHeartbeat = true //是否开启心跳
const HeartbeatInterval = 10  // 心跳间隔时间，单位秒

type Message struct {
	Method     string      `json:"method"`
	FromUserId string      `json:"from_user_id"`
	Data       interface{} `json:"data"` // map[string]interface{}
}

// 保存用户与连接的映射
var (
	UserIdToConn = make(map[uint]*websocket.Conn) // userID -> conn
	ConnToUserId = make(map[*websocket.Conn]uint) // userID -> conn
	ConnHeart    = make(map[uint]int64)           // userID -> lastHeartbeatTime
	ClientsMu    sync.RWMutex                     // 读写锁，防止并发问题
)

// WsHandler 处理 WebSocket 连接
func WsHandler(c *gin.Context) {
	userID := c.Value("UserId").(uint)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("升级WebSocket出错:", err)
		return
	}
	defer func() {
		CloseConn(conn)
		fmt.Printf("用户%d已断开链接...\n", userID)
	}()

	//链接信息维护
	AddConnection(userID, conn)
	if IsStartHeartbeat {
		//心跳机制
		go KeepAlive(conn)
	}

	fmt.Printf("用户%d已连接成功\n", userID)

	// 循环接收消息
	for {
		mt, msg, err1 := conn.ReadMessage()
		if err1 != nil {
			fmt.Println("读取websocket消息错误:", err1)
			return
		}
		//记录最后一次活跃的时间
		ConnHeart[userID] = time.Now().Unix()
		fmt.Printf("收到用户%d的消息: %s\n", userID, msg)
		// 回复客户端
		conn.WriteMessage(mt, []byte("服务器已收到: "+string(msg)))
	}
}

func CloseConn(conn *websocket.Conn) {
	ClientsMu.Lock()
	defer ClientsMu.Unlock()
	conn.Close()
}
