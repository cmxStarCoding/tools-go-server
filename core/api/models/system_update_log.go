package models

import (
	"gorm.io/gorm"
)

// SystemUpdateLogModel 系统更新日志模型
type SystemUpdateLogModel struct {
	ID                uint           `gorm:"column:id" json:"id"`
	Version           string         `gorm:"column:version" json:"version" `
	VersionName       string         `gorm:"column:version_name" json:"version_name" `
	Content           string         `gorm:"column:content" json:"content"`
	MacDownloadUrl    string         `gorm:"column:mac_download_url" json:"mac_download_url"`
	WindowDownloadUrl string         `gorm:"column:window_download_url" json:"window_download_url"`
	CreatedAt         TimeNormal     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         TimeNormal     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (SystemUpdateLogModel) TableName() string {
	return "t_system_update_log"
}
