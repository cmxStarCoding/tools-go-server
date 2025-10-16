package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"
const timezone = "Asia/Shanghai"

type TimeNormal time.Time

// MarshalJSON 用于将 TimeNormal 类型转换为 JSON 字节数组的方法。在该方法中，时间被格式化为指定的字符串格式，并以 JSON 字符串的形式返回。
func (t TimeNormal) MarshalJSON() ([]byte, error) {
	// 检查时间是否为零值

	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON 用于从 JSON 字节数组解析并设置 TimeNormal 类型的值。在该方法中，通过 time.ParseInLocation 解析 JSON 字符串，并将其转换为 TimeNormal 类型
func (t *TimeNormal) UnmarshalJSON(data []byte) (err error) {

	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = TimeNormal(now)
	return
}

// 将 TimeNormal 类型转换为字符串的方法。在该方法中，时间被格式化为指定的字符串格式。
func (t TimeNormal) String() string {
	return time.Time(t).Format(timeFormat)
}

// 将 TimeNormal 类型的时间转换为本地时区的方法。
func (t TimeNormal) local() time.Time {
	loc, _ := time.LoadLocation(timezone)
	return time.Time(t).In(loc)
}

// Value 实现 database/sql/driver.Valuer 接口的方法，用于将 TimeNormal 类型的时间转换为数据库驱动的值。如果时间为零值，则返回 nil。
func (t TimeNormal) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
	//return ti.Format(timeFormat), nil
}

// Scan 实现 database/sql.Scanner 接口的方法，用于从数据库扫描操作中读取时间值，并设置为 TimeNormal 类型。
func (t *TimeNormal) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = TimeNormal(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// ParseTimeString 将年月日时分秒的字符串转换为 TimeNormal 类型。
func ParseTimeString(s string) (TimeNormal, error) {
	t, err := time.Parse(timeFormat, s)
	if err != nil {
		return TimeNormal(time.Time{}), err
	}
	return TimeNormal(t), nil
}
