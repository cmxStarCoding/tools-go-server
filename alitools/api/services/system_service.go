package services

import (
	"errors"
	"fmt"
	"github.com/Masterminds/semver"
	"gorm.io/gorm"
	"journey/alitools/api/validator"
	models2 "journey/alitools/models"
	"journey/common/database"
)

type SystemService struct {
}

func (s SystemService) FeedBack(requestData *validator.FeedbackRequest, UserId uint) (string, error) {

	feedback := &models2.FeedbackModel{}
	feedback.UserId = UserId
	feedback.Content = requestData.Content
	feedback.ContractPhone = requestData.ContractPhone
	database.DB.Save(feedback)
	return "ok", nil
}

func (s SystemService) SystemUpdateLog(requestData *validator.GetUpdateLogRequest) (map[string]interface{}, error) {

	var mapResult = make(map[string]interface{})

	var sliceUpdateLog []models2.SystemUpdateLogModel

	var total int64
	database.DB.Model(&models2.SystemUpdateLogModel{}).Count(&total)
	database.DB.Limit(int(requestData.Limit)).Offset((int(requestData.Page) - 1) * int(requestData.Limit)).Find(&sliceUpdateLog)

	mapResult["total"] = total
	mapResult["list"] = sliceUpdateLog

	return mapResult, nil
}

func (s SystemService) CheckSystemUpdate(requestData *validator.CheckSystemUpdateRequest) (map[string]any, error) {

	systemUpdate := &models2.SystemUpdateLogModel{}
	result := database.DB.Last(&systemUpdate)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("不存在版本")
	}
	version1, _ := semver.NewVersion(requestData.ClientVersion)
	version2, _ := semver.NewVersion(systemUpdate.Version)

	returnMap := make(map[string]any)
	returnMap["is_exist_new_version"] = 0

	// 比较版本号
	if version1.LessThan(version2) {
		returnMap["is_exist_new_version"] = 1
		returnMap["new_version"] = systemUpdate
	}
	return returnMap, nil
}

func (s SystemService) CurrentLatestVersion() (*models2.SystemUpdateLogModel, error) {

	systemUpdate := &models2.SystemUpdateLogModel{}
	result := database.DB.Last(&systemUpdate)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("不存在版本")
	}

	return systemUpdate, nil
}
