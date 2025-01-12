package models

import (
	"gorm.io/gorm"
)

type UserResetPasswordModel struct {
	ID        uint           `gorm:"column:id" json:"id"`
	Account   string         `gorm:"column:account" json:"account"`
	UseEmail  string         `gorm:"column:use_email" json:"use_email"`
	Code      string         `gorm:"column:code" json:"code"`
	CreatedAt TimeNormal     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt TimeNormal     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (UserResetPasswordModel) TableName() string {
	return "t_user_reset_password"
}
