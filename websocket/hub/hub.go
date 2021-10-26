package hub

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"websocket/data"
)

var H = Hub{
	Conn:    make(map[*Connection]bool),
	U:       make(chan *Connection),
	B:       make(chan []byte),
	Receive: make(chan *Connection),
}

type Hub struct {
	Conn    map[*Connection]bool
	B       chan []byte
	Receive chan *Connection
	U       chan *Connection
}

func (h *Hub) Run() {
	for {
		select {
		case ch := <-h.Receive:
			H.Conn[ch] = true
			ch.Data.Ip = ch.Ws.RemoteAddr().String()
			ch.Data.Type = "handshake"
			ch.Data.UserList = User_list
			data_b, _ := json.Marshal(ch.Data)
			ch.Sc <- data_b
		case ch := <-h.U:
			if _, ok := h.Conn[ch]; ok {
				delete(h.Conn, ch)
				close(ch.Sc)
			}
		case data := <-h.B:
			for c := range h.Conn {
				select {
				case c.Sc <- data:
				default:
					delete(h.Conn, c)
					close(c.Sc)
				}
			}
		}
	}
}

var wu = &websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Task(w http.ResponseWriter, r *http.Request) {
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &Connection{
		Ws:   ws,
		Sc:   make(chan []byte, 256),
		Data: &data.Data{},
	}
	H.Receive <- c
	go c.Writer()
	c.Reader()
	defer func() {
		c.Data.Type = "logout"
		User_list = Del(User_list, c.Data.User)
		c.Data.UserList = User_list
		c.Data.Content = c.Data.User
		data_b, _ := json.Marshal(c.Data)
		H.B <- data_b
		H.Receive <- c
	}()
}

type Connection struct {
	Ws   *websocket.Conn
	Sc   chan []byte
	Data *data.Data
}

func (c *Connection) Writer() {
	for message := range c.Sc {
		c.Ws.WriteMessage(websocket.TextMessage, message)
	}
	c.Ws.Close()
}

var User_list []string

func (c *Connection) Reader() {
	for {
		_, message, err := c.Ws.ReadMessage()
		if err != nil {
			H.Receive <- c
			break
		}
		json.Unmarshal(message, &c.Data)
		switch c.Data.Type {
		case "login":
			c.Data.User = c.Data.Content
			c.Data.From = c.Data.User
			User_list = append(User_list, c.Data.User)
			c.Data.UserList = User_list
			data_b, _ := json.Marshal(c.Data)
			H.B <- data_b
		case "user":
			c.Data.Type = "user"
			data_b, _ := json.Marshal(c.Data)
			H.B <- data_b
		case "logout":
			c.Data.Type = "logout"
			User_list = Del(User_list, c.Data.User)
		default:
			fmt.Println("========default================")
		}
	}
}

func Del(slice []string, user string) []string {
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
