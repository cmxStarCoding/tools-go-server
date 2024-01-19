package models

import (
	"gorm.io/gorm"
)

// SystemUpdateLogModel 系统更新日志模型
type SystemUpdateLogModel struct {
	ID        uint           `gorm:"column:id" json:"id"`
	Version   string         `gorm:"column:version" json:"version" `
	Content   string         `gorm:"column:content" json:"content"`
	CreatedAt TimeNormal     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt TimeNormal     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (SystemUpdateLogModel) TableName() string {
	return "t_system_update_log"
}
