package dao

import (
	"errors"
	"github.com/google/uuid"
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

// CreateGroup 创建群组
func CreateGroup(username string, group model.Group) error {
	User, err := GetUserByUsername(username)
	if err != nil {
		return errors.New("用户不存在")
	}

	group.UserID = User.ID
	group.Uuid = uuid.New().String()

	if err = DB.Model(model.Group{}).Create(&group).Error; err != nil {
		return errors.New("创建群组失败")
	}

	groupMember := model.GroupMember{
		GroupID:  group.ID,
		UserID:   User.ID,
		UserUuid: User.Uuid,
		NickName: User.NickName,
		Avatar:   User.Avatar,
		Mute:     0,
		Role:     2,
	}

	if err = DB.Model(model.GroupMember{}).Create(&groupMember).Error; err != nil {
		return errors.New("创建群组失败")
	}

	return nil
}

// GetGroupMembers 获取群成员
func GetGroupMembers(groupUuId string) ([]model.GroupMember, error) {
	var group model.Group
	err := DB.Model(&model.Group{}).Where("uuid = ?", groupUuId).First(&group).Error
	if err != nil {
		return nil, errors.New("群组不存在")
	}

	var members []model.GroupMember
	err = DB.Model(&model.GroupMember{}).Where("group_id = ?", group.ID).Find(&members).Error
	if err != nil {
		return nil, errors.New("查询群成员失败")
	}
	return members, nil
}

// JoinGroup 用户加入群组
func JoinGroup(username string, groupUuid string) error {
	// 获取用户信息
	User, err := GetUserByUsername(username)
	if err != nil {
		return errors.New("用户不存在")
	}

	// 查找群组
	var group model.Group
	err = DB.Model(&model.Group{}).Where("uuid = ?", groupUuid).First(&group).Error
	if err != nil || group.ID <= 0 {
		return errors.New("群组不存在")
	}

	Group := model.Group{
		UserID:    User.ID,
		Uuid:      groupUuid,
		GroupName: group.GroupName,
		Notice:    group.Notice,
		Avatar:    group.Avatar,
	}

	// 检查用户是否已经在群组中
	var existingMember model.GroupMember
	err = DB.Model(&model.GroupMember{}).
		Where("group_id = ? AND user_id = ?", group.ID, User.ID).
		First(&existingMember).Error
	if err == nil {
		return errors.New("用户已在群组中，不能重复加入")
	}

	// 创建群组成员记录
	groupMember := model.GroupMember{
		GroupID:  group.ID,
		UserID:   User.ID,
		UserUuid: User.Uuid,
		NickName: User.NickName,
		Avatar:   User.Avatar,
		Mute:     0,
		Role:     0,
	}

	// 插入群组成员记录
	err = DB.Model(&model.GroupMember{}).Create(&groupMember).Error
	err = DB.Model(&model.Group{}).Create(&Group).Error
	if err != nil {
		return errors.New("加入群组失败")
	}

	return nil
}

// QuitGroup 用户退出群组
func QuitGroup(username string, groupUuid string) error {
	// 获取用户信息
	User, err := GetUserByUsername(username)
	if err != nil {
		return errors.New("用户不存在")
	}

	// 查找群组
	var group model.Group
	err = DB.Model(&model.Group{}).Where("uuid = ?", groupUuid).First(&group).Error
	if err != nil || group.ID <= 0 {
		return errors.New("群组不存在")
	}

	// 删除群组成员记录
	err = DB.Model(&model.GroupMember{}).
		Where("group_id = ? AND user_id = ?", group.ID, User.ID).
		Delete(&model.GroupMember{}).Error
	if err != nil {
		return errors.New("退出群组失败")
	}

	// 删除群组记录
	err = DB.Model(&model.Group{}).
		Where("uuid = ? And user_id = ?", groupUuid, User.ID).
		Delete(&model.Group{}).Error
	if err != nil {
		return errors.New("退出群组失败")
	}

	return nil
}
