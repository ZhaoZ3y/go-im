package dao

import (
	"errors"
	"goim/model"
	"goim/model/model_json"
	"gorm.io/gorm"
)

// GetMessagesBySolo 获取私聊消息列表
func GetMessagesBySolo(username string, uuid string) ([]model_json.Message, error) {
	User, _ := GetUserByUsername(username)
	ToUser, err := GetUserByUuid(uuid)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("该用户不存在")
	}

	var messages []model_json.Message
	err = DB.Raw(`SELECT m.id, m.from_user_id, m.to_user_id, m.content, m.content_type, m.url, m.created_at, u.username 
    	AS from_username, u.avatar, to_user.username AS to_username  FROM messages 
    	AS m LEFT JOIN users AS u ON m.from_user_id = u.id LEFT JOIN users AS to_user ON m.to_user_id = to_user.id 
    	WHERE from_user_id IN (?, ?) AND to_user_id IN (?, ?)`,
		User.ID, ToUser.ID, User.ID, ToUser.ID).Scan(&messages).Error
	if err != nil {
		return nil, errors.New("查询消息列表失败")
	}

	return messages, nil
}

// GetMessagesByGroup 获取群聊消息列表
func GetMessagesByGroup(groupUuid string) ([]model_json.Message, error) {
	var group model.Group
	err := DB.Model(&model.Group{}).Where("uuid = ?", groupUuid).First(&group).Error
	if err != nil {
		return nil, errors.New("群组不存在")
	}

	var messages []model_json.Message
	err = DB.Raw(`SELECT m.id, m.from_user_id, m.to_user_id, m.content, m.content_type, m.url, m.created_at, u.username 
    	AS from_username, u.avatar FROM messages AS m LEFT JOIN users 
    	AS u ON m.from_user_id = u.id WHERE m.message_type = 2 AND m.to_user_id = ?`,
		group.ID).Scan(&messages).Error
	if err != nil {
		return nil, errors.New("查询消息列表失败")
	}

	return messages, nil
}

// SaveMessageBySolo 保存消息
func SaveMessageBySolo(username string, ToUuid string, content string, contentType int, messageType int, url string) error {
	User, _ := GetUserByUsername(username)
	ToUser, err := GetUserByUuid(ToUuid)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该用户不存在")
	}

	message := model.Message{
		FromUserID:  User.ID,
		ToUserID:    ToUser.ID,
		Content:     content,
		ContentType: int64(contentType),
		MessageType: int64(messageType),
		Url:         url,
	}

	err = DB.Create(&message).Error
	if err != nil {
		return errors.New("保存消息失败")
	}

	return nil
}

// SaveMessageByGroup 保存群聊消息
func SaveMessageByGroup(username string, groupUuid string, content string, contentType int, messageType int, url string) error {
	User, _ := GetUserByUsername(username)
	var group model.Group
	err := DB.Model(&model.Group{}).Where("uuid = ?", groupUuid).First(&group).Error
	if err != nil {
		return errors.New("群组不存在")
	}

	message := model.Message{
		FromUserID:  User.ID,
		ToUserID:    group.ID,
		Content:     content,
		ContentType: int64(contentType),
		MessageType: int64(messageType),
		Url:         url,
	}

	err = DB.Create(&message).Error
	if err != nil {
		return errors.New("保存消息失败")
	}

	return nil
}
