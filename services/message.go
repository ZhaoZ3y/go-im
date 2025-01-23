package services

import (
	"goim/dao"
	"goim/model"
	"goim/model/model_json"
	constant "goim/utils/const"
	"goim/utils/protocol"
)

// GetMessages 获取用户群组
func GetMessages(username string, message model.MessageReq) ([]model_json.Message, error) {
	if message.MessageType == constant.MESSAGE_TYPE_USER {
		return dao.GetMessagesBySolo(username, message.Uuid)
	} else if message.MessageType == constant.MESSAGE_TYPE_GROUP {
		return dao.GetMessagesByGroup(message.Uuid)
	}
	return nil, nil
}

// SaveMessage 保存消息
func SaveMessage(username string, message protocol.Message) error {
	if message.MessageType == constant.MESSAGE_TYPE_USER {
		return dao.SaveMessageBySolo(username, message.To, message.Content, int(message.ContentType), int(message.MessageType), message.Url)
	} else if message.MessageType == constant.MESSAGE_TYPE_GROUP {
		return dao.SaveMessageByGroup(username, message.To, message.Content, int(message.ContentType), int(message.MessageType), message.Url)
	}
	return nil
}
