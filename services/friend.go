package services

import (
	"errors"
	"goim/dao"
	"goim/model/model_json"
	"sort"
)

// GetFriendList 获取好友列表
func GetFriendList(username string) ([]model_json.Friend, error) {
	// 查询好友列表
	friends, err := dao.GetFriendsList(username)
	if err != nil {
		return nil, errors.New("查询好友列表失败")
	}

	var friendsJson []model_json.Friend
	for _, friend := range friends {
		friendJson := model_json.Friend{
			NickName: friend.NickName,
			Username: friend.UserName,
			Uuid:     friend.Uuid,
			Avatar:   friend.Avatar,
		}
		friendsJson = append(friendsJson, friendJson)
	}

	// 按照nickname的首字母从小到大排序
	sort.Slice(friendsJson, func(i, j int) bool {
		return friendsJson[i].NickName < friendsJson[j].NickName
	})

	return friendsJson, nil
}

// AddFriend 添加好友
func AddFriend(username, friendName string) error {
	err := dao.AddFriend(username, friendName)
	if err != nil {
		if err.Error() == "已经是好友了" {
			return errors.New("已经是好友了")
		}
		return errors.New("添加好友失败")
	}
	return nil
}

// DeleteFriend 删除好友
func DeleteFriend(username, friendName string) error {
	err := dao.DeleteFriend(username, friendName)
	if err != nil {
		return errors.New("删除好友失败")
	}
	return nil
}
