package services

import (
	"tools/common/database"
	"tools/core/api/models"
	"tools/core/api/validator/category"
)

// CategoryService 用户服务
type CategoryService struct{}

func (s CategoryService) GetCategoryList(requestData *category.GetCategoryListRequest) ([]models.CategoryModel, error) {
	db := database.DB
	var categoryList []models.CategoryModel
	db.Where("pid = 0").Limit(int(requestData.Limit)).Offset((int(requestData.Page) - 1) * int(requestData.Limit)).
		Preload("Children").Find(&categoryList)
	return categoryList, nil
}
