package models

import (
	"gorm.io/gorm"
	"time"
)

// CategoryModel UserModel 用户模型
type CategoryModel struct {
	gorm.Model
	ID          uint           `gorm:"column:id" json:"id"`
	Name        string         `gorm:"column:name" json:"name"`
	Description string         `gorm:"column:description" json:"description" `
	Pid         uint           `gorm:"column:pid" json:"pid"`
	Status      uint           `gorm:"column:status" json:"status"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`

	Children []CategoryModel `gorm:"foreignKey:pid" json:"children"`
	Tools    []ToolsModel    `gorm:"foreignKey:CategoryId" json:"tools"`
}

func (CategoryModel) TableName() string {
	return "t_category"
}
