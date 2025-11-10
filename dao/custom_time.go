package dao

import (
	"fmt"
	"time"
)

// CustomTime 用于自定义时间格式化输出
type CustomTime time.Time

// MarshalJSON 实现 JSON 序列化为 "Y-m-d H:i:s"
func (t CustomTime) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte(`""`), nil
	}
	formatted := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// ToTime 返回原生 time.Time 对象（可选辅助方法）
func (t CustomTime) ToTime() time.Time {
	return time.Time(t)
}
