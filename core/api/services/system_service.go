package services

import (
	"errors"
	"fmt"
	"github.com/Masterminds/semver"
	"gorm.io/gorm"
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

func (s SystemService) CheckSystemUpdate(requestData *system.CheckSystemUpdateRequest) (map[string]any, error) {

	systemUpdate := &models.SystemUpdateLogModel{}
	result := database.DB.Last(&systemUpdate)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("不存在版本")
	}
	version1, _ := semver.NewVersion(requestData.ClientVersion)
	version2, _ := semver.NewVersion(systemUpdate.Version)

	returnMap := make(map[string]any)
	returnMap["is_exist_version"] = 0

	// 比较版本号
	if version1.LessThan(version2) {
		returnMap["is_exist_version"] = 1
		returnMap["version"] = systemUpdate
	}
	return returnMap, nil
}
