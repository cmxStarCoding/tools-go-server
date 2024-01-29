package services

import (
	"encoding/json"
	"fmt"
	"tools/common/database"
	"tools/common/utils"
	"tools/core/api/models"
	"tools/core/api/validator/pic"
	"tools/core/api/validator/user"
)

type UserTaskLogService struct{}

func (s UserTaskLogService) GetUserTaskLogList(requestData *user.GetTaskLogListRequest, UserId uint) (map[string]interface{}, error) {

	type ExtendedUserTaskLog struct {
		models.UserTaskLogModel
		StatusText string `gorm:"-" json:"status_text"`
	}

	var mapList []ExtendedUserTaskLog
	query := database.DB.Where("user_id = ?", UserId)
	if requestData.Status > 0 {
		query.Where("status = ?", requestData.Status)
	}
	var total int64
	query.Model(&models.UserTaskLogModel{}).Count(&total)
	query.Limit(int(requestData.Limit)).Offset((int(requestData.Page) - 1) * int(requestData.Limit)).Preload("Tools").Preload("User").Find(&mapList)

	var result = make(map[string]interface{})
	result["total"] = total
	result["list"] = mapList

	if len(mapList) > 0 {
		for i := range mapList {
			if mapList[i].Status == 1 {
				mapList[i].StatusText = "等待中"
			} else if mapList[i].Status == 2 {
				mapList[i].StatusText = "执行中"
			} else if mapList[i].Status == 3 {
				mapList[i].StatusText = "执行成功"
			} else if mapList[i].Status == 4 {
				mapList[i].StatusText = "执行失败"
			}
		}
	}
	return result, nil
}

func (s UserTaskLogService) CreateTask(strategyId uint, toolsMark string, UserId uint) (*models.UserPicPasteStrategyModel, string, error) {
	tools := &models.ToolsModel{}
	database.DB.Where("mark = ?", toolsMark).First(tools)

	userPicStrategy := &models.UserPicPasteStrategyModel{}
	database.DB.Where("id = ?", strategyId).First(userPicStrategy)

	taskIdString := utils.GenerateUniqueRandomString()
	requestData, _ := json.Marshal(userPicStrategy)
	UserTaskLog := models.UserTaskLogModel{
		ToolId:      tools.ID,
		UserId:      UserId,
		TaskId:      taskIdString,
		Status:      1,
		RequestData: string(requestData),

		RequestResult: "{}",
	}
	result := database.DB.Create(&UserTaskLog)
	if result.Error != nil {
		// 处理错误  错误原因 是 result.Error.Error()
		return nil, "", fmt.Errorf("创建任务失败")
	}

	//转map
	jsonStr, _ := json.Marshal(UserTaskLog)
	var InitMap = make(map[string]any)
	if err3 := json.Unmarshal(jsonStr, &InitMap); err3 != nil {
		return nil, "", fmt.Errorf("map失败")
	}

	// 解析request_data为JSON对象
	var requestDataObj = make(map[string]any)
	if err1 := json.Unmarshal([]byte(UserTaskLog.RequestData), &requestDataObj); err1 != nil {
		return nil, "", fmt.Errorf("map失败")
	}

	InitMap["request_data"] = requestDataObj
	return userPicStrategy, taskIdString, nil
}

func (s UserTaskLogService) EditTaskStatus(request *pic.NotifyRequest) {
	if request.Status <= 0 {
		return
	}
	userTaskLog := &models.UserTaskLogModel{}
	resultError := database.DB.Where("task_id = ?", request.BatchNo).First(userTaskLog)
	if resultError.Error != nil {
		return
	}
	jsonStr, _ := json.Marshal(request)

	if request.Status == 1 {
		userTaskLog.Status = 3
	} else {
		userTaskLog.Status = 4
	}

	userTaskLog.RequestResult = string(jsonStr)
	database.DB.Save(userTaskLog)
	//添加一条使用记录
	userUseLog := &models.UserUseLogModel{}
	userUseLog.ToolId = userTaskLog.ToolId
	userUseLog.UserId = userTaskLog.UserId
	database.DB.Save(userUseLog)
}
