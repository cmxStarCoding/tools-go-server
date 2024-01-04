package services

import (
	"tools/common/database"
	"tools/core/api/models"
	"tools/core/api/validator/vip"
)

type VipLevelService struct{}

func (s VipLevelService) GetVipLevelList(requestData *vip.GetVipLevelListRequest) ([]models.VipLevelModel, error) {

	var sliceVip []models.VipLevelModel
	database.DB.Where("status = ?", 1).Limit(int(requestData.Limit)).Offset((int(requestData.Page) - 1) * int(requestData.Limit)).Find(&sliceVip)

	return sliceVip, nil
}
