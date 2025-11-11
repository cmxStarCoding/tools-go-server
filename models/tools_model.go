package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// 定义常量
const (
	PicPasteMark = "pic_paste_mark"
)

type TToolsModel struct {
	ID          uint           `gorm:"column:id" json:"id"`
	Mark        string         `gorm:"column:mark" json:"mark"`
	Name        string         `gorm:"column:name" json:"name"`
	Logo        string         `gorm:"column:logo" json:"logo"`
	Description string         `gorm:"column:description" json:"description"`
	CategoryId  uint           `gorm:"column:category_id" json:"category_id"`
	IsRecommend uint           `gorm:"column:is_recommend" json:"is_recommend"`
	Router      string         `gorm:"column:router" json:"router"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// MarshalJSON 自定义时间格式输出（Y-m-d H:i:s）
func (u TToolsModel) MarshalJSON() ([]byte, error) {
	type Alias TToolsModel
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

func (TToolsModel) TableName() string {
	return "t_tools"
}
