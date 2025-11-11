package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type TUserResetPasswordModel struct {
	ID        uint           `gorm:"column:id" json:"id"`
	Account   string         `gorm:"column:account" json:"account"`
	UseEmail  string         `gorm:"column:use_email" json:"use_email"`
	Code      string         `gorm:"column:code" json:"code"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (u TUserResetPasswordModel) MarshalJSON() ([]byte, error) {
	type Alias TUserResetPasswordModel
	return json.Marshal(&struct {
		VipExpireTime string `json:"vip_expire_time"`
		CreatedAt     string `json:"created_at"`
		UpdatedAt     string `json:"updated_at"`
		DeletedAt     string `json:"deleted_at"`
		*Alias
	}{
		CreatedAt: formatTime(u.CreatedAt),
		UpdatedAt: formatTime(u.UpdatedAt),
		DeletedAt: formatDeletedAt(u.DeletedAt),
		Alias:     (*Alias)(&u),
	})
}

func (TUserResetPasswordModel) TableName() string {
	return "t_user_reset_password"
}
