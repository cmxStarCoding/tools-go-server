package services

import (
	"tools/common/database"
	"tools/core/api/models"
	"tools/core/api/validator/user"
)

type UserUseLogService struct{}

func (s UserUseLogService) UserUseLogList(requestData *user.GetUserUseLogListRequest, UserId uint) ([]models.UserUseLogModel, error) {

	var sliceLog []models.UserUseLogModel
	database.DB.Where("user_id = ?", UserId).Preload("User").Limit(int(requestData.Limit)).Offset((int(requestData.Page) - 1) * int(requestData.Limit)).Find(&sliceLog)

	return sliceLog, nil
}
