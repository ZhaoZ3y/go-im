package dao

import (
	"goim/model"
)

// IsExistedUser 判断用户是否存在
func IsExistedUser(username string) (bool, error) {
	var userCount int64
	result := DB.Model(&model.User{}).Where("user_name = ?", username).Count(&userCount)
	if result.Error != nil {
		return false, result.Error // 数据库查询失败
	}

	// 用户存在则返回 true，否则返回 false
	if userCount > 0 {
		return true, nil
	}

	return false, nil
}

// CreateUser 创建用户
func CreateUser(user *model.User) error {
	result := DB.Create(user)
	return result.Error
}
