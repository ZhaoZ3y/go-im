package services

import (
	"errors"
	"goim/dao"
	"goim/model"
	"goim/model/model_json"
	"goim/utils/crypto"
)

// UpdateUser 更新用户信息
func UpdateUser(username string, user model.User) error {
	err := dao.UpdateUser(username, user)
	if err != nil {
		return errors.New("更新用户信息失败")
	}
	return nil
}

// GetUserByUsername 获取用户信息
func GetUserByUsername(username string) (model_json.User, error) {
	user, err := dao.GetUserByUsername(username)
	if err != nil {
		return model_json.User{}, errors.New("查询用户失败")
	}

	userJson := model_json.User{
		ID:       int64(user.ID),
		Uuid:     user.Uuid,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Avatar:   user.Avatar,
	}

	return userJson, nil
}

// SearchUser 搜索用户
func SearchUser(username string) ([]model_json.User, error) {
	users, err := dao.SearchUser(username)
	if err != nil {
		return nil, errors.New("查询用户失败")
	}

	var usersJson []model_json.User
	for _, user := range users {
		userJson := model_json.User{
			ID:       int64(user.ID),
			Uuid:     user.Uuid,
			UserName: user.UserName,
			NickName: user.NickName,
			Email:    user.Email,
			Avatar:   user.Avatar,
		}
		usersJson = append(usersJson, userJson)
	}

	return usersJson, nil
}

// ChangePassword 修改密码
func ChangePassword(username string, oldPassword string, newPassword string) error {
	HashNewPassword := crypto.CryptoPwd(newPassword)
	HashOldPassword := crypto.CryptoPwd(oldPassword)

	_, err := dao.CheckUser(username, HashOldPassword)
	if err != nil {
		return errors.New("账号或密码错误")
	}

	if err := dao.ChangePassword(username, HashNewPassword); err != nil {
		return errors.New("修改密码失败")
	}
	return nil

}

// ChangeAvatar 修改头像
func ChangeAvatar(username string, avatar string) error {
	err := dao.ChangeAvatar(username, avatar)
	if err != nil {
		return errors.New("修改头像失败")
	}
	return nil
}
