package models

type UserUseLogModel struct {
	ID        uint       `gorm:"column:id" json:"id"`
	UserId    uint       `gorm:"column:user_id" json:"user_id"`
	ToolId    uint       `gorm:"column:tool_id" json:"tool_id"`
	CreatedAt TimeNormal `gorm:"column:created_at" json:"created_at"`
	UpdatedAt TimeNormal `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt TimeNormal `gorm:"column:deleted_at" json:"deleted_at"`
	User      UserModel  `json:"user"`
	Tool      ToolsModel `json:"tool"`
}

func (UserUseLogModel) TableName() string {
	return "t_user_use_log"
}
