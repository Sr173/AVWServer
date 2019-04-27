package protocol

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

func WeosocketHandler(w http.ResponseWriter, r*http.Request) {
	u := websocket.Upgrader{ReadBufferSize: 1, WriteBufferSize: 10000}
	c,err := u.Upgrade(w,r,nil)
	if err != nil  {
		return
	}
	fmt.Print(c.RemoteAddr().String())
	fmt.Println("ping")

	for {
		if true {
			c.WriteMessage(websocket.TextMessage,[]byte("请求就绪"))
			time.Sleep(time.Millisecond * 100)
		} else {
			c.Close()
			return
		}
	}
}
