package dao

import (
	"errors"
	"goim/model"
)

// GetGroups 获取用户群组
func GetGroups(userId uint) ([]model.Group, error) {
	var groups []model.Group
	err := DB.Model(&model.Group{}).Where("user_id = ?", userId).Find(&groups).Error
	if err != nil {
		return nil, errors.New("查询用户群组失败")
	}
	return groups, nil
}
