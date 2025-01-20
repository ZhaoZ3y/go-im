package services

import (
	"errors"
	"goim/dao"
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
