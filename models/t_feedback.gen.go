package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// TFeedback 意见反馈模型
type TFeedback struct {
	ID            uint           `gorm:"column:id" json:"id"`
	UserId        uint           `gorm:"column:user_id" json:"user_id"`
	ContractPhone string         `gorm:"column:contract_phone" json:"contract_phone" `
	Content       string         `gorm:"column:content" json:"content"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (u TFeedback) MarshalJSON() ([]byte, error) {
	type Alias TFeedback
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

func (TFeedback) TableName() string {
	return "t_feedback"
}
