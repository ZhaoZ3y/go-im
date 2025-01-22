package services

import (
	"errors"
	"goim/dao"
	"goim/model"
	"goim/model/model_json"
	"sort"
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

	// 按照name的首字母从小到大排序
	sort.Slice(groupsJson, func(i, j int) bool {
		return groupsJson[i].Name < groupsJson[j].Name
	})

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
func GetGroupMembers(groupUuId string) ([]model.GroupMember, error) {
	members, err := dao.GetGroupMembers(groupUuId)
	if err != nil {
		return nil, errors.New("查询群成员失败")
	}
	return members, nil
}

// JoinGroup 加入群组
func JoinGroup(username string, groupUuid string) error {
	err := dao.JoinGroup(username, groupUuid)
	if err != nil {
		if err.Error() == "用户已在群组中，不能重复加入" {
			return errors.New("用户已存在")
		}
		return errors.New("加入群组失败")
	}
	return nil
}

// QuitGroup 退出群组
func QuitGroup(username string, groupUuid string) error {
	err := dao.QuitGroup(username, groupUuid)
	if err != nil {
		if err.Error() == "群组不存在" {
			return errors.New("群组不存在")
		}
		return errors.New("退出群组失败")
	}
	return nil
}
