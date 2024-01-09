package services

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
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

func (s UserPicPasteStrategyService) editUserPicPasteStrategyService(requestData *pic.SaveUserPicPasteStrategyRequest, UserId uint) (models.UserPicPasteStrategyModel, error) {
	userPicPasteStrategy := &models.UserPicPasteStrategyModel{}
	resultErr := database.DB.Where("id = ?", requestData.ID).Where("user_id = ?", UserId).First(&userPicPasteStrategy)
	if resultErr != nil && errors.Is(resultErr.Error, gorm.ErrRecordNotFound) {
		return models.UserPicPasteStrategyModel{}, fmt.Errorf("记录数据不存在")
	}
	userPicPasteStrategy.OriginalImageUrl = requestData.OriginalImageUrl
	userPicPasteStrategy.StickImgUrl = requestData.StickImgUrl
	userPicPasteStrategy.X = requestData.X
	userPicPasteStrategy.Y = requestData.Y
	userPicPasteStrategy.R = requestData.R
	userPicPasteStrategy.Type = requestData.Type
	userPicPasteStrategy.Multiple = requestData.Multiple
	userPicPasteStrategy.IsSquare = requestData.IsSquare
	userPicPasteStrategy.SideLength = requestData.SideLength
	database.DB.Save(&userPicPasteStrategy)
	return *userPicPasteStrategy, nil
}

func (s UserPicPasteStrategyService) createUserPicPasteStrategyService(requestData *pic.SaveUserPicPasteStrategyRequest, UserId uint) (models.UserPicPasteStrategyModel, error) {
	userPicPasteStrategy := models.UserPicPasteStrategyModel{
		UserId:           UserId,
		OriginalImageUrl: requestData.OriginalImageUrl,
		StickImgUrl:      requestData.StickImgUrl,
		X:                requestData.X,
		Y:                requestData.Y,
		R:                requestData.R,
		Type:             requestData.Type,
		Multiple:         requestData.Multiple,
		IsSquare:         requestData.IsSquare,
		SideLength:       requestData.SideLength,
	}
	result := database.DB.Create(&userPicPasteStrategy)
	if result.Error != nil {
		return models.UserPicPasteStrategyModel{}, fmt.Errorf("创建贴图策略失败")
	}
	return userPicPasteStrategy, nil
}

func (s UserPicPasteStrategyService) SaveUserPicPasteStrategy(requestData *pic.SaveUserPicPasteStrategyRequest, UserId uint) (models.UserPicPasteStrategyModel, error) {
	//更新
	if requestData.ID > 0 {
		return s.editUserPicPasteStrategyService(requestData, UserId)
	}
	return s.createUserPicPasteStrategyService(requestData, UserId)

}

func (s UserPicPasteStrategyService) DeleteUserPicPasteStrategy(id uint, UserId uint) (string, error) {
	userPicPasteStrategy := &models.UserPicPasteStrategyModel{}
	resultErr := database.DB.Where("id = ?", id).Where("user_id", UserId).First(&userPicPasteStrategy)

	if resultErr != nil && errors.Is(resultErr.Error, gorm.ErrRecordNotFound) {
		return "", fmt.Errorf("记录数据不存在")
	}
	deleteResult := database.DB.Delete(&userPicPasteStrategy)
	if deleteResult.RowsAffected > 0 {
		fmt.Printf("删除了%d行\n", deleteResult.RowsAffected)
	}
	return "ok", nil
}
