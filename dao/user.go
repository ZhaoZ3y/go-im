package dao

import (
	"errors"
	"goim/model"
	"gorm.io/gorm"
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
	err := DB.Create(user).Error
	return err
}

// GetUserByUsername 获取用户信息
func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := DB.Where("user_name = ?", username).First(&user).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在")
	} else if err != nil {
		return nil, errors.New("查询用户失败")
	}
	return &user, nil
}

// CheckUser 核实账号密码
func CheckUser(username, password string) (model.User, error) {
	var user model.User
	err := DB.Where("user_name = ? AND pass_word = ?", username, password).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
