package models

import (
	"gorm.io/gorm"
)

// VipLevelModel 用户模型
type VipLevelModel struct {
	gorm.Model
	ID        uint           `gorm:"column:id" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	Price     float32        `gorm:"column:price" json:"price"`
	Status    uint           `gorm:"column:status" json:"status"`
	CreatedAt TimeNormal     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt TimeNormal     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (VipLevelModel) TableName() string {
	return "t_vip_level"
}
