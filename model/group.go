package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	GroupName string `gorm:"type:varchar(150);unique_index;not null" comment:"群名称"`
	UserID    uint   `gorm:"index:idx_groups_user_id;" comment:"群主ID"`
	Notice    string `gorm:"type:text;default null" comment:"群公告"`
	Uuid      string `gorm:"type:varchar(150);uniqueIndex:idx_uuid;not null" comment:"群唯一标识"`
	Avatar    string `gorm:"type:varchar(255);default null" comment:"群头像"`
}

func (g *Group) TableName() string {
	return "groups"
}

type GroupMember struct {
	gorm.Model
	GroupID  uint   `gorm:"index:idx_group_members_group_id;" comment:"群ID"`
	UserID   uint   `gorm:"index:idx_group_members_user_id;" comment:"用户ID"`
	NickName string `gorm:"type:varchar(255);not null" comment:"昵称"`
	Mute     *int16 `gorm:"type:smallint;default 0" comment:"是否禁言：0否，1是"`
}

func (m *GroupMember) TableName() string {
	return "group_members"
}
