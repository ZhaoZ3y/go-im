package socket

import (
	"encoding/base64"
	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"goim/config"
	"goim/dao"
	"goim/services"
	constant "goim/utils/const"
	"goim/utils/file"
	"goim/utils/protocol"
	"os"
	"strings"
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

// Run 运行服务端
func (s *Server) Run() {
	for {
		select {
		case conn := <-s.Register:
			s.mutex.Lock()
			s.Clients[conn.Name] = conn
			s.mutex.Unlock()
			msg := &protocol.Message{
				From:    "System",
				To:      conn.Name,
				Content: conn.Name + "上线",
			}
			protoMsg, _ := proto.Marshal(msg)
			conn.Send <- protoMsg

		case conn := <-s.Unregister:
			s.mutex.Lock()
			if _, ok := s.Clients[conn.Name]; ok {
				close(conn.Send)
				delete(s.Clients, conn.Name)
			}
			s.mutex.Unlock()

		case message := <-s.Broadcast:
			msg := &protocol.Message{}
			proto.Unmarshal(message, msg)

			if msg.To != "" {
				// 一般消息，比如文本消息，视频文件消息等
				if msg.ContentType >= constant.TEXT && msg.ContentType <= constant.VIDEO {
					// 保存消息只会在存在socket的一个端上进行保存，防止分布式部署后，消息重复问题
					_, exist := s.Clients[msg.From]
					if exist {
						SaveMessage(msg)
					}

					if msg.MessageType == constant.MESSAGE_TYPE_USER {
						clients, ok := s.Clients[msg.To]
						if ok {
							msgByte, err := proto.Marshal(msg)
							if err != nil {
								return
							}
							clients.Send <- msgByte
						}
					} else if msg.MessageType == constant.MESSAGE_TYPE_GROUP {
						sendMessageToGroup(msg, s)
					}
				} else {
					// 语音电话，视频电话等，仅支持单人聊天，不支持群聊
					// 不保存文件，直接进行转发
					client, ok := s.Clients[msg.To]
					if ok {
						client.Send <- message
					}
				}
			} else {
				// 无对应接受人员进行广播
				for _, conn := range s.Clients {
					select {
					case conn.Send <- message:
					default:
						close(conn.Send)
						delete(s.Clients, conn.Name)
					}
				}
			}
		}
	}
}

// SaveMessage 保存消息，如果是文本消息直接保存，如果是文件，语音等消息，保存文件后，保存对应的文件路径
func SaveMessage(msg *protocol.Message) {
	// 如果上传的是base64文件，需要先保存文件，然后再保存文件路径
	if msg.ContentType == constant.FILE {
		// 保存文件
		// 保存文件路径
		url := uuid.New().String() + ".png"
		index := strings.Index(msg.Content, "base64")
		index += 7

		// 保存文件
		content := msg.Content
		content = content[index:]

		dataBuffer, dateErr := base64.StdEncoding.DecodeString(content)
		if dateErr != nil {
			return
		}
		err := os.WriteFile(config.GetConfig().Path.FilePath+url, dataBuffer, 0666)
		if err != nil {
			return
		}
		msg.Url = url
		msg.Content = ""
	} else if msg.ContentType == constant.IMAGE {
		// 普通二进制文件，直接保存文件路径
		fileSuffix := file.GetFileType(msg.File)
		nullStr := ""
		if fileSuffix == nullStr {
			fileSuffix = strings.ToLower(msg.FileSuffix)
		}

		contentType := file.GetContentTypeBySuffix(fileSuffix)
		url := uuid.New().String() + "." + fileSuffix
		err := os.WriteFile(config.GetConfig().Path.FilePath+url, msg.File, 0666)
		if err != nil {
			return
		}
		msg.Url = url
		msg.File = nil
		msg.ContentType = contentType
	}

	err := services.SaveMessage(msg.FromUsername, *msg)
	if err != nil {
		return
	}
}

// SendMessageToGroup 发送消息到群组
func sendMessageToGroup(msg *protocol.Message, s *Server) {
	// 发送给群组的消息，查找该群所有的用户进行发送
	users, err := dao.GetGroupMembers(msg.To)
	if err != nil {
		return
	}

	for _, user := range users {
		if user.UserUuid == msg.From {
			continue
		}

		s.mutex.Lock()
		clients, ok := s.Clients[user.UserUuid]
		s.mutex.Unlock()
		if !ok {
			continue
		}

		fromUserDetail, _ := dao.GetUserByUuid(msg.From)
		msgSend := protocol.Message{
			Avatar:       fromUserDetail.Avatar,
			FromUsername: msg.FromUsername,
			From:         msg.To,
			To:           msg.From,
			Content:      msg.Content,
			ContentType:  msg.ContentType,
			MessageType:  msg.MessageType,
			Url:          msg.Url,
		}

		msgByte, err := proto.Marshal(&msgSend)
		if err != nil {
			return
		}
		clients.Send <- msgByte
	}
}
