package services

import (
	"errors"
	"goim/dao"
	"goim/model"
	"goim/model/model_json"
)

// GetGroup 获取用户群组`
func GetGroup(userId uint) ([]model_json.Group, error) {
	groups, err := dao.GetGroups(userId)
	if err != nil {
		return nil, errors.New("查询用户群组失败")
	}

	var groupsJson []model_json.Group
	for _, group := range groups {
		groupJson := model_json.Group{
			ID:        int64(group.ID),
			Uuid:      group.Uuid,
			Name:      group.GroupName,
			Avatar:    group.Avatar,
			Notice:    group.Notice,
			CreatedAt: group.CreatedAt,
		}
		groupsJson = append(groupsJson, groupJson)
	}

	return groupsJson, nil
}

// CreateGroup 创建群组
func CreateGroup(username string, group model_json.Group) error {
	Group := model.Group{
		GroupName: group.Name,
		Notice:    group.Notice,
		Avatar:    group.Avatar,
	}

	err := dao.CreateGroup(username, Group)
	if err != nil {
		return errors.New("创建群组失败")
	}
	return nil
}

// GetGroupMembers 获取群成员
func GetGroupMembers(groupId uint) ([]model.GroupMember, error) {
	members, err := dao.GetGroupMembers(groupId)
	if err != nil {
		return nil, errors.New("查询群成员失败")
	}
	return members, nil
}

// JoinGroup 加入群组
func JoinGroup(username string, groupUuid string) error {
	err := dao.JoinGroup(username, groupUuid)
	if err != nil {
		return errors.New("加入群组失败")
	}
	return nil
}
