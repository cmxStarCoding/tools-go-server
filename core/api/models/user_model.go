package models

import (
	"gorm.io/gorm"
	"time"
)

// UserModel 用户模型
type UserModel struct {
	gorm.Model
	ID            uint           `gorm:"column:id" json:"id"`
	Phone         string         `gorm:"column:phone" json:"phone"`
	Password      string         `gorm:"column:password <-:create" json:"password" `
	Nickname      string         `gorm:"column:nickname" json:"nickname"`
	AvatarUrl     string         `gorm:"column:avatar_url" json:"avatar_url"`
	VipLevelId    uint16         `gorm:"column:vip_level_id" json:"vip_level_id"`
	VipExpireTime time.Time      `gorm:"column:vip_expire_time" json:"vip_expire_time"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (UserModel) TableName() string {
	return "t_user"
}
