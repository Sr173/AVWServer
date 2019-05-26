package protocol

import (
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

var Conn_map = make(map[string]chan string)

func WeosocketHandler(w http.ResponseWriter, r*http.Request) {
	u := websocket.Upgrader{ReadBufferSize: 1, WriteBufferSize: 10000}
	c,err := u.Upgrade(w,r,nil)

	if err != nil  {
		return
	}
	remote := c.RemoteAddr().String()
	remote = remote[0:strings.Index(remote,":")]
	message_list := make(chan string)
	Conn_map[remote] = message_list
	for {
		message := <- message_list
		c.WriteMessage(websocket.TextMessage,[]byte(message))
	}
}
