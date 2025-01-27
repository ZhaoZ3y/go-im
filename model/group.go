package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	GroupName string `gorm:"type:varchar(255);not null" comment:"群名称"`
	UserID    uint   `gorm:"index:idx_groups_user_id;" comment:"群主ID"`
	Notice    string `gorm:"type:text;default null" comment:"群公告"`
	Uuid      string `gorm:"type:varchar(150);not null" comment:"群唯一标识"`
	Avatar    string `gorm:"type:varchar(255);default null" comment:"群头像"`
}

func (g *Group) TableName() string {
	return "groups"
}

type GroupMember struct {
	gorm.Model
	GroupID  uint   `gorm:"index:idx_group_members_group_id;" comment:"群ID"`
	UserID   uint   `gorm:"index:idx_group_members_user_id;" comment:"用户ID"`
	UserUuid string `gorm:"type:varchar(150);not null" comment:"用户唯一标识"`
	NickName string `gorm:"type:varchar(255);not null" comment:"昵称"`
	Avatar   string `gorm:"type:varchar(255);default null" comment:"头像"`
	Mute     int    `gorm:"type:smallint;default 0" comment:"是否禁言：0否，1是"`
	Role     int    `gorm:"type:smallint;default 0" comment:"角色：0普通成员，1管理员，2群主"`
}

func (m *GroupMember) TableName() string {
	return "group_members"
}
