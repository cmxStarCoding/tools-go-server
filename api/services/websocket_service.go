package services

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type Data struct {
	Ip       string   `json:"ip"`
	User     string   `json:"user"`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"user_list"`
}

type WebsocketService struct{}

type connection struct {
	ws   *websocket.Conn
	sc   chan []byte
	data *Data
}

type hub struct {
	c map[*connection]bool
	b chan []byte
	r chan *connection
	u chan *connection
}

var h = hub{
	c: make(map[*connection]bool), // 存储当前所有连接的客户端
	u: make(chan *connection),     // 用于接收断开连接的通道
	b: make(chan []byte),          // 用于广播消息的通道
	r: make(chan *connection),     //新链接
}

func (s WebsocketService) Run() {
	fmt.Println("进入run逻辑")
	for {
		select {
		case c := <-h.r: //新链接
			h.c[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			//当前链接的握手消息
			fmt.Println("新键握手消息成功")
			c.sc <- data_b
		case c := <-h.u: //断开连接 删除并关闭链接
			if _, ok := h.c[c]; ok {
				delete(h.c, c)
				close(c.sc)
			}
		case data := <-h.b: //广播消息
			for c := range h.c { //所有已经建立链接的websocket用户
				select {
				case c.sc <- data: // 向单个链接websocket用户发送消息
				default: //无法向 c.sc 通道发送数据，则删除并关闭链接
					delete(h.c, c)
					close(c.sc)
				}
			}
		}
	}
}

var wu = &websocket.Upgrader{ReadBufferSize: 512,
	WriteBufferSize: 512, CheckOrigin: func(r *http.Request) bool { return true }}

func (s WebsocketService) MyWs(w http.ResponseWriter, r *http.Request) {
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	fmt.Println("WebSocket connection established")
	c := &connection{sc: make(chan []byte, 256), ws: ws, data: &Data{}}
	h.r <- c

	// 将当前用户的消息发送到当前的websocket链接
	go c.writer()
	//读取用户发送的消息，做对应处理
	c.reader()

	defer func() {
		c.data.Type = "logout"
		user_list = del(user_list, c.data.User)
		c.data.UserList = user_list
		c.data.Content = c.data.User
		data_b, _ := json.Marshal(c.data)
		h.b <- data_b
		h.r <- c
	}()
}

func (c *connection) writer() {
	//当前websocket用户消息通道
	for message := range c.sc {
		//向当前链接发送消息
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

var user_list = []string{}

func (c *connection) reader() {
	fmt.Println("用户邓丽了11")
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			h.r <- c
			break
		}
		json.Unmarshal(message, &c.data)
		switch c.data.Type {
		case "login":
			c.data.User = c.data.Content
			c.data.From = c.data.User
			user_list = append(user_list, c.data.User)
			c.data.UserList = user_list
			fmt.Println("用户邓丽了")
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		case "user":
			c.data.Type = "user"
			data_b, _ := json.Marshal(c.data)
			fmt.Println("收到用户消息")
			h.b <- data_b
		case "logout":
			c.data.Type = "logout"
			user_list = del(user_list, c.data.User)
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
			h.r <- c
		default:
			fmt.Print("========default================")
		}
	}
}

func del(slice []string, user string) []string {
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}
	var n_slice = []string{}
	for i := range slice {
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			n_slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	fmt.Println(n_slice)
	return n_slice
}
