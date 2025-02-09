package services

import (
	"journey/alitools/api/validator"
	"journey/alitools/models"
	"journey/common/database"
)

// CategoryService 用户服务
type CategoryService struct{}

func (s CategoryService) GetCategoryList(requestData *validator.GetCategoryListRequest) ([]models.CategoryModel, error) {
	db := database.DB
	var categoryList []models.CategoryModel
	db.Where("pid = 0").Limit(int(requestData.Limit)).Offset((int(requestData.Page) - 1) * int(requestData.Limit)).
		Preload("Children").Find(&categoryList)
	return categoryList, nil
}

func (s CategoryService) GetCategoryToolsList(requestData *validator.GetCategoryToolsRequest) ([]models.CategoryModel, error) {
	db := database.DB
	var categoryList []models.CategoryModel
	db.Where("pid = 0").Limit(int(requestData.Limit)).Offset((int(requestData.Page) - 1) * int(requestData.Limit)).
		Preload("Tools").Find(&categoryList)
	return categoryList, nil
}
