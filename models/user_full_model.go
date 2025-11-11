package models

import (
	"gorm.io/gorm"
)

// TUser 用户模型
type UserFullModel struct {
	ID            uint           `gorm:"column:id" json:"id"`
	Phone         string         `gorm:"column:phone" json:"phone"`
	Account       string         `gorm:"column:account" json:"account"`
	Password      string         `gorm:"password" json:"password" `
	Nickname      string         `gorm:"column:nickname" json:"nickname"`
	AvatarUrl     string         `gorm:"column:avatar_url" json:"avatar_url"`
	VipLevelId    uint16         `gorm:"column:vip_level_id" json:"vip_level_id"`
	VipExpireTime TimeNormal     `gorm:"column:vip_expire_time" json:"vip_expire_time"`
	CreatedAt     TimeNormal     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     TimeNormal     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (UserFullModel) TableName() string {
	return "t_user"
}
