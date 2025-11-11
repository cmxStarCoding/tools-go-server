package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// formatTime 格式化普通时间
func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

// formatDeletedAt 格式化 gorm.DeletedAt
func formatDeletedAt(d gorm.DeletedAt) string {
	if d.Valid {
		return d.Time.Format("2006-01-02 15:04:05")
	}
	return ""
}

func formatNullTime(t sql.NullTime) string {
	if t.Valid {
		return t.Time.Format("2006-01-02 15:04:05")
	}
	return ""
}
