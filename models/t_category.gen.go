package models

import (
	"gorm.io/gorm"
)

// Category TUser 用户模型
type Category struct {
	ID          uint           `gorm:"column:id" json:"id"`
	Name        string         `gorm:"column:name" json:"name"`
	Description string         `gorm:"column:description" json:"description" `
	Pid         uint           `gorm:"column:pid" json:"pid"`
	Status      uint           `gorm:"column:status" json:"status"`
	CreatedAt   TimeNormal     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   TimeNormal     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`

	Children []Category `gorm:"foreignKey:pid" json:"children"`
	Tools    []TTools   `gorm:"foreignKey:CategoryId" json:"tools"`
}

func (Category) TableName() string {
	return "t_category"
}
