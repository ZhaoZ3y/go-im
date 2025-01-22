package dao

import (
	"errors"
	"goim/model"
)

// GetFriendsList 获取好友列表
func GetFriendsList(username string) ([]model.User, error) {
	// 获取用户信息
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 查询好友列表
	var friends []model.User
	// 查询好友列表
	err = DB.Raw(`
		SELECT u.user_name AS user_name, u.uuid, u.avatar, u.nick_name
		FROM user_friends AS uf
		JOIN users AS u ON uf.friend_id = u.id
		WHERE uf.user_id = ? AND uf.deleted_at IS NULL
	`, user.ID).Scan(&friends).Error

	if err != nil {
		return nil, errors.New("查询好友列表失败")
	}

	return friends, nil
}

// AddFriend 添加好友
func AddFriend(username, friendName string) error {
	User, _ := GetUserByUsername(username)
	Friend, err := GetUserByUsername(friendName)
	if err != nil {
		return errors.New("用户不存在")
	}

	var FriendCount int64
	DB.Model(&model.UserFriend{}).Where("user_id = ? AND friend_id = ?", User.ID, Friend.ID).Count(&FriendCount)
	if FriendCount > 0 {
		return errors.New("已经是好友了")
	}

	friend := model.UserFriend{
		UserID:   User.ID,
		FriendID: Friend.ID,
	}

	err = DB.Model(&model.UserFriend{}).Create(&friend).Error
	if err != nil {
		return errors.New("添加好友失败")
	}

	return nil
}

// DeleteFriend 删除好友
func DeleteFriend(username, friendName string) error {
	User, _ := GetUserByUsername(username)
	Friend, err := GetUserByUsername(friendName)
	if err != nil {
		return errors.New("用户不存在")
	}

	var FriendCount int64
	DB.Model(&model.UserFriend{}).Where("user_id = ? AND friend_id = ?", User.ID, Friend.ID).Count(&FriendCount)
	if FriendCount == 0 {
		return errors.New("不是好友关系")
	}

	err = DB.Model(&model.UserFriend{}).
		Where("user_id = ? AND friend_id = ?", User.ID, Friend.ID).
		Delete(&model.UserFriend{}).Error
	if err != nil {
		return errors.New("删除好友失败")
	}

	return nil
}
