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

func (s UserPicPasteStrategyService) GetUserPicPasteStrategyList(requestData *pic.GetUserPicPasteStrategyListRequest, UserId uint) (map[string]interface{}, error) {
	var total int64
	database.DB.Model(&models.UserPicPasteStrategyModel{}).Count(&total)

	type ReturnResult struct {
		models.UserPicPasteStrategyModel
		ShapeText string `gorm:"-" json:"shape_text"`
		TypeText  string `gorm:"-" json:"type_text"`
	}
	var list []ReturnResult
	resultErr := database.DB.Where("user_id = ?", UserId).Preload("User").Limit(int(requestData.Limit)).Offset((int(requestData.Page) - 1) * int(requestData.Limit)).Find(&list)
	if resultErr.Error != nil {
		return nil, resultErr.Error
	}
	var marResult = make(map[string]interface{})
	marResult["total"] = total

	if len(list) > 0 {
		for i := range list {
			if list[i].BcShape == 1 {
				list[i].ShapeText = "圆形"
			} else if list[i].BcShape == 2 {
				list[i].ShapeText = "方形"
			} else {
				list[i].ShapeText = "无背景区域"
			}
			multiple := fmt.Sprintf("%.2f", list[i].Multiple)
			if list[i].Type == 1 {
				list[i].TypeText = "放大" + multiple + "倍"
			} else if list[i].Type == 2 {
				list[i].TypeText = "缩小" + multiple + "倍"
			} else {
				list[i].TypeText = "无缩放"
			}

		}
	}
	marResult["list"] = list
	return marResult, nil
}

func (s UserPicPasteStrategyService) editUserPicPasteStrategyService(requestData *pic.SaveUserPicPasteStrategyRequest, UserId uint) (models.UserPicPasteStrategyModel, error) {
	userPicPasteStrategy := &models.UserPicPasteStrategyModel{}
	resultErr := database.DB.Where("id = ?", requestData.ID).Where("user_id = ?", UserId).First(&userPicPasteStrategy)
	if resultErr != nil && errors.Is(resultErr.Error, gorm.ErrRecordNotFound) {
		return models.UserPicPasteStrategyModel{}, fmt.Errorf("记录数据不存在")
	}
	userPicPasteStrategy.Name = requestData.Name
	userPicPasteStrategy.OriginalImageUrl = requestData.OriginalImageUrl
	userPicPasteStrategy.StickImgUrl = requestData.StickImgUrl
	userPicPasteStrategy.X = requestData.X
	userPicPasteStrategy.Y = requestData.Y
	userPicPasteStrategy.R = requestData.R
	userPicPasteStrategy.Type = requestData.Type
	userPicPasteStrategy.Multiple = requestData.Multiple
	userPicPasteStrategy.BcShape = requestData.BcShape
	userPicPasteStrategy.BcColor = requestData.BcColor

	userPicPasteStrategy.SideLength = requestData.SideLength
	database.DB.Save(&userPicPasteStrategy)
	return *userPicPasteStrategy, nil
}

func (s UserPicPasteStrategyService) createUserPicPasteStrategyService(requestData *pic.SaveUserPicPasteStrategyRequest, UserId uint) (models.UserPicPasteStrategyModel, error) {
	userPicPasteStrategy := models.UserPicPasteStrategyModel{
		UserId:           UserId,
		Name:             requestData.Name,
		OriginalImageUrl: requestData.OriginalImageUrl,
		StickImgUrl:      requestData.StickImgUrl,
		X:                requestData.X,
		Y:                requestData.Y,
		R:                requestData.R,
		Type:             requestData.Type,
		Multiple:         requestData.Multiple,
		BcShape:          requestData.BcShape,
		BcColor:          requestData.BcColor,
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

	userPicPasteStrategy := &models.UserPicPasteStrategyModel{}
	resultErr := database.DB.Where("name = ?", requestData.Name).Where("user_id", UserId).First(&userPicPasteStrategy)
	if resultErr.Error != nil {
		if errors.Is(resultErr.Error, gorm.ErrRecordNotFound) {
			return s.createUserPicPasteStrategyService(requestData, UserId)
		}
		return models.UserPicPasteStrategyModel{}, fmt.Errorf(resultErr.Error.Error())
	}
	return models.UserPicPasteStrategyModel{}, fmt.Errorf("策略名称已存在")
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
