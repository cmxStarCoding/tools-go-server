package services

import (
	"tools/common/database"
	"tools/core/api/models"
	"tools/core/api/validator/pic"
)

type UserPicPasteStrategyService struct {
}

func (s UserPicPasteStrategyService) GetUserPicPasteStrategyList(requestData *pic.GetUserPicPasteStrategyListRequest, UserId uint) ([]models.UserPicPasteStrategyModel, error) {

	var slice []models.UserPicPasteStrategyModel
	resultErr := database.DB.Where("user_id = ?", UserId).Preload("User").Limit(int(requestData.Limit)).Offset((int(requestData.Page) - 1) * int(requestData.Limit)).Find(&slice)
	if resultErr.Error != nil {
		return nil, resultErr.Error
	}
	return slice, nil
}

func (s UserPicPasteStrategyService) SaveUserPicPasteStrategy() {

}

func (s UserPicPasteStrategyService) DeleteUserPicPasteStrategy() {

}
