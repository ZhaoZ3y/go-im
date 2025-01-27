package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"goim/infra/socket"
	"net/http"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// RunSocket 运行socket
func RunSocket(c *gin.Context) {
	user := c.Query("user")
	if user == "" {
		return
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &socket.Client{
		Name: user,
		Conn: ws,
		Send: make(chan []byte),
	}

	socket.WsSever.Register <- client
	go client.Read()
	go client.Write()
}
