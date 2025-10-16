package models

import (
	"gorm.io/gorm"
)

// FeedbackModel 意见反馈模型
type FeedbackModel struct {
	ID            uint           `gorm:"column:id" json:"id"`
	UserId        uint           `gorm:"column:user_id" json:"user_id"`
	ContractPhone string         `gorm:"column:contract_phone" json:"contract_phone" `
	Content       string         `gorm:"column:content" json:"content"`
	CreatedAt     TimeNormal     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     TimeNormal     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (FeedbackModel) TableName() string {
	return "t_feedback"
}
