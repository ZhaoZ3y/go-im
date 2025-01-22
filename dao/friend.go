package dao

import "goim/model"

// GetFriendsList 获取好友列表
func GetFriendsList() ([]model.User, error) {
	var users []model.User
	err := DB.Table("users").
		Joins("left join user_friends on users.id = user_friends.friend_id").
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
