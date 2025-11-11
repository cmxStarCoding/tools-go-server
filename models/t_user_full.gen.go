package models

import (
	"database/sql"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// TUserFull 用户模型
type TUserFull struct {
	ID            uint           `gorm:"column:id" json:"id"`
	Phone         string         `gorm:"column:phone" json:"phone"`
	Account       string         `gorm:"column:account" json:"account"`
	Password      string         `gorm:"password" json:"password" `
	Nickname      string         `gorm:"column:nickname" json:"nickname"`
	AvatarUrl     string         `gorm:"column:avatar_url" json:"avatar_url"`
	VipLevelId    uint16         `gorm:"column:vip_level_id" json:"vip_level_id"`
	VipExpireTime sql.NullTime   `gorm:"column:vip_expire_time" json:"vip_expire_time"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// MarshalJSON 自定义时间格式输出（Y-m-d H:i:s）
func (u TUserFull) MarshalJSON() ([]byte, error) {
	type Alias TUserFull
	return json.Marshal(&struct {
		VipExpireTime string `json:"vip_expire_time"`
		CreatedAt     string `json:"created_at"`
		UpdatedAt     string `json:"updated_at"`
		DeletedAt     string `json:"deleted_at"`
		*Alias
	}{
		VipExpireTime: formatNullTime(u.VipExpireTime),
		CreatedAt:     formatTime(u.CreatedAt),
		UpdatedAt:     formatTime(u.UpdatedAt),
		DeletedAt:     formatDeletedAt(u.DeletedAt),
		Alias:         (*Alias)(&u),
	})
}

func (TUserFull) TableName() string {
	return "t_user"
}
