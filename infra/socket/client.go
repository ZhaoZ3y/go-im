package socket

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"goim/config"
	"goim/infra/mq"
	constant "goim/utils/const"
	"goim/utils/protocol"
)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
	Name string
}

func (c *Client) Read() {
	defer func() {
		WsSever.Unregister <- c
		c.Conn.Close()
	}()

	for {
		c.Conn.PongHandler()
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			WsSever.Unregister <- c
			c.Conn.Close()
			break
		}

		msg := &protocol.Message{}
		err = proto.Unmarshal(message, msg)

		// pong
		if msg.Type == constant.HEAT_BEAT {
			pong := &protocol.Message{
				Content: constant.PONG,
				Type:    constant.HEAT_BEAT,
			}
			protoMsg, _ := proto.Marshal(pong)
			c.Conn.WriteMessage(websocket.BinaryMessage, protoMsg)
		} else {
			if config.GetConfig().MsgChannel.ChannelType == constant.KAFKA {
				// 发送消息到kafka
				mq.SendMessage(message)
			} else {
				WsSever.Broadcast <- message
			}
		}
	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		err := c.Conn.WriteMessage(websocket.BinaryMessage, message)
		if err != nil {
			break
		}

	}
}
