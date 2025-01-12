package models

import (
	"gorm.io/gorm"
	"time"
)

type UserPicPasteStrategyModel struct {
	ID               uint           `gorm:"column:id" json:"id"`
	UserId           uint           `gorm:"column:user_id" json:"user_id"`
	Name             string         `gorm:"column:name" json:"name"`
	OriginalImageUrl string         `gorm:"column:original_image_url" json:"original_image_url"`
	StickImgUrl      string         `gorm:"column:stick_img_url" json:"stick_img_url"`
	X                uint           `gorm:"column:x" json:"x"`
	Y                uint           `gorm:"column:y" json:"y"`
	R                uint           `gorm:"column:r" json:"r"`
	Type             uint           `gorm:"column:type" json:"type"`
	Multiple         float32        `gorm:"column:multiple" json:"multiple"`
	BcShape          uint           `gorm:"column:bc_shape" json:"bc_shape"`
	BcColor          string         `gorm:"column:bc_color" json:"bc_color"`
	SideLength       uint           `gorm:"column:side_length" json:"side_length"`
	CreatedAt        time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	User             UserModel      `json:"user"`
}

func (UserPicPasteStrategyModel) TableName() string {
	return "t_user_pic_paste_strategy"
}
