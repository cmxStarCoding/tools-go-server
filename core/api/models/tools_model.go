package models

import (
	"gorm.io/gorm"
	"time"
)

// 定义常量
const (
	PicPasteMark = "pic_paste_mark"
)

type ToolsModel struct {
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

func (ToolsModel) TableName() string {
	return "t_tools"
}
