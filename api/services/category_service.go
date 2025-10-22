package services

import (
	"fmt"
	"gorm.io/gorm"
	"journey/api/validator"
	"journey/common/database"
	"journey/models"
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

	sqlStr := db.Model(&models.CategoryModel{}).
		Where("pid = 0").Limit(int(requestData.Limit)).Offset((int(requestData.Page) - 1) * int(requestData.Limit)).
		ToSQL(func(tx *gorm.DB) *gorm.DB {
			return tx.Find(&categoryList)
		})
	fmt.Println("执行sql", sqlStr)

	db.Where("pid = 0").Limit(int(requestData.Limit)).Offset((int(requestData.Page) - 1) * int(requestData.Limit)).
		Preload("Tools").Find(&categoryList)
	return categoryList, nil
}
