package models

import (
	"gorm.io/gorm"
	"time"
)

type UserUseLogModel struct {
	gorm.Model
	ID        uint           `gorm:"column:id" json:"id"`
	UserId    uint           `gorm:"column:user_id" json:"user_id"`
	ToolId    uint           `gorm:"column:tool_id" json:"tool_id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	User      UserModel      `json:"user"`
}

func (UserUseLogModel) TableName() string {
	return "t_user_use_log"
}
