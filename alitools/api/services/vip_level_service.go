package services

import (
	"journey/alitools/api/validator"
	"journey/alitools/models"
	"journey/common/database"
)

type VipLevelService struct{}

func (s VipLevelService) GetVipLevelList(requestData *validator.GetVipLevelListRequest) ([]models.VipLevelModel, error) {

	var sliceVip []models.VipLevelModel
	database.DB.Where("status = ?", 1).Limit(int(requestData.Limit)).Offset((int(requestData.Page) - 1) * int(requestData.Limit)).Find(&sliceVip)

	return sliceVip, nil
}
