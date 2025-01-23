package socket

import (
	"sync"
)

var WsSever = NewSever()

// Server 服务端
type Server struct {
	Clients    map[string]*Client
	mutex      *sync.Mutex
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

// NewSever 创建服务端
func NewSever() *Server {
	return &Server{
		Clients:    make(map[string]*Client),
		mutex:      &sync.Mutex{},
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

// ConsumerKafkaMessage 消费消息
func ConsumerKafkaMessage(msg []byte) {
	WsSever.Broadcast <- msg
}
