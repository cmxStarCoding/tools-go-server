package services

import (
	"tools/common/database"
	"tools/core/api/models"
	"tools/core/api/validator/system"
)

type SystemService struct {
}

func (s SystemService) FeedBack(requestData *system.FeedbackRequest, UserId uint) (string, error) {

	feedback := &models.FeedbackModel{}
	feedback.UserId = UserId
	feedback.Content = requestData.Content
	feedback.ContractPhone = requestData.ContractPhone
	database.DB.Save(feedback)
	return "ok", nil
}

func (s SystemService) SystemUpdateLog(requestData *system.GetUpdateLogRequest) ([]models.SystemUpdateLogModel, error) {

	var sliceUpdateLog []models.SystemUpdateLogModel

	database.DB.Limit(int(requestData.Limit)).Offset((int(requestData.Page) - 1) * int(requestData.Limit)).Find(&sliceUpdateLog)

	return sliceUpdateLog, nil
}
