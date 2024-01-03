package services

import (
	"encoding/json"
	"fmt"
	"tools/common/database"
	"tools/common/utils"
	"tools/core/api/models"
	"tools/core/api/validator/pic"
)

type UserTaskLogService struct{}

func (s UserTaskLogService) CreateTask(request *pic.Request, toolsMark string, UserId uint) (map[string]any, error) {
	tools := &models.ToolsModel{}
	database.DB.Where("mark = ?", toolsMark).First(tools)
	taskIdString := utils.GenerateUniqueRandomString()
	requestData, _ := json.Marshal(request)
	UserTaskLog := models.UserTaskLogModel{
		ToolId:        tools.ID,
		UserId:        UserId,
		TaskId:        taskIdString,
		RequestData:   string(requestData),
		RequestResult: "{}",
	}
	result := database.DB.Create(&UserTaskLog)
	if result.Error != nil {
		// 处理错误  错误原因 是 result.Error.Error()
		return nil, fmt.Errorf("创建任务失败")
	}

	//转map
	jsonStr, _ := json.Marshal(UserTaskLog)
	var InitMap = make(map[string]any)
	if err3 := json.Unmarshal(jsonStr, &InitMap); err3 != nil {
		return nil, fmt.Errorf("map失败")
	}

	// 解析request_data为JSON对象
	var requestDataObj = make(map[string]any)
	if err1 := json.Unmarshal([]byte(UserTaskLog.RequestData), &requestDataObj); err1 != nil {
		return nil, fmt.Errorf("map失败")
	}

	InitMap["request_data"] = requestDataObj
	return InitMap, nil
}

func (s UserTaskLogService) EditTaskStatus(request *pic.NotifyRequest) {

	if request.Status == 1 {
		userTaskLog := &models.UserTaskLogModel{}
		database.DB.Where("task_id = ?", request.BatchNo).First(userTaskLog)
		userTaskLog.Status = 1
		database.DB.Save(userTaskLog)
	}
}
