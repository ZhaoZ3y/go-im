package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uuid     string `gorm:"type:varchar(150);uniqueIndex:idx_uuid;not null" comment:"用户唯一标识"`
	UserName string `gorm:"type:varchar(150);unique_index;not null" comment:"用户名"`
	NickName string `gorm:"type:varchar(255);not null" comment:"昵称"`
	Email    string `gorm:"type:varchar(255);default null" comment:"邮箱"`
	PassWord string `gorm:"type:varchar(255);not null" comment:"密码"`
	Avatar   string `gorm:"type:varchar(255);not null" comment:"头像"`
}

func (u *User) TableName() string {
	return "users"
}

type UserFriend struct {
	gorm.Model
	UserID   uint `gorm:"index:idx_user_friends_user_id;" comment:"用户ID"`
	FriendID uint `gorm:"index:idx_user_friends_friend_id;" comment:"好友ID"`
}

func (f *UserFriend) TableName() string {
	return "user_friends"
}
