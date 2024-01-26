package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type UserTaskJSON json.RawMessage

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
//func (j *UserTaskJSON) Scan(value interface{}) error {
//	bytes, ok := value.([]byte)
//	if !ok {
//		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
//	}
//
//	result := json.RawMessage{}
//	err := json.Unmarshal(bytes, &result)
//	*j = UserTaskJSON(result)
//	return err
//}
//
//// Value 实现 driver.Valuer 接口，Value 返回 json value
//func (j UserTaskJSON) Value() (driver.Value, error) {
//	if len(j) == 0 {
//		return nil, nil
//	}
//	return json.RawMessage(j).MarshalJSON()
//}

type UserTaskLogModel struct {
	ID            uint           `gorm:"column:id" json:"id"`
	ToolId        uint           `gorm:"column:tool_id" json:"tool_id"`
	UserId        uint           `gorm:"column:user_id" json:"user_id"`
	TaskId        string         `gorm:"column:task_id" json:"task_id"`
	Status        uint           `gorm:"column:status" json:"status"`
	RequestData   string         `gorm:"column:request_data" json:"request_data"`
	RequestResult string         `gorm:"column:request_result" json:"request_result"`
	UserFailMsg   string         `gorm:"column:user_fail_msg" json:"user_fail_msg"`
	SystemFailMsg string         `gorm:"column:system_fail_msg" json:"system_fail_msg"`
	CreatedAt     TimeNormal     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     TimeNormal     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	User          UserModel      `json:"user"`
	Tools         ToolsModel     `gorm:"foreignKey:ToolId" json:"tools"`
}

func (UserTaskLogModel) TableName() string {
	return "t_user_task_log"
}
