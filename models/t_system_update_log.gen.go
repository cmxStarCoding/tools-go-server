package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// TSystemUpdateLog 系统更新日志模型
type TSystemUpdateLog struct {
	ID                uint           `gorm:"column:id" json:"id"`
	Version           string         `gorm:"column:version" json:"version" `
	VersionName       string         `gorm:"column:version_name" json:"version_name" `
	Content           string         `gorm:"column:content" json:"content"`
	IntervalPeriod    uint           `gorm:"column:interval_period" json:"interval_period"`
	MacDownloadUrl    string         `gorm:"column:mac_download_url" json:"mac_download_url"`
	WindowDownloadUrl string         `gorm:"column:window_download_url" json:"window_download_url"`
	CreatedAt         time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// MarshalJSON 自定义时间格式输出（Y-m-d H:i:s）
func (u TSystemUpdateLog) MarshalJSON() ([]byte, error) {
	type Alias TSystemUpdateLog
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
func (TSystemUpdateLog) TableName() string {
	return "t_system_update_log"
}
