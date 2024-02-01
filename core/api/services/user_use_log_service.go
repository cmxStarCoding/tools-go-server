package services

import (
	"tools/common/database"
	"tools/core/api/models"
	"tools/core/api/validator/user"
)

type UserUseLogService struct{}

func (s UserUseLogService) UserUseLogList(requestData *user.GetUserUseLogListRequest, UserId uint) (map[string]interface{}, error) {

	var mapResult = make(map[string]interface{})

	var sliceLog []models.UserUseLogModel
	var total int64
	database.DB.Model(&models.UserUseLogModel{}).Where("user_id = ?", UserId).Count(&total)
	database.DB.Where("user_id = ?", UserId).Preload("User").Preload("Tool").Limit(int(requestData.Limit)).Offset((int(requestData.Page) - 1) * int(requestData.Limit)).Find(&sliceLog)

	mapResult["total"] = total
	mapResult["list"] = sliceLog

	return mapResult, nil
}
