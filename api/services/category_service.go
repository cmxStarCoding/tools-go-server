package services

import (
	"fmt"
	"journey/api/validator"
	"journey/common/database"
	"journey/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (s CategoryService) GetCategoryToolsList(ctx *gin.Context, requestData *validator.GetCategoryToolsRequest) ([]models.CategoryModel, error) {
	db := database.DB
	var categoryList []models.CategoryModel

	commonQuery := db.Model(&models.CategoryModel{}).
		Where("pid = 0")

	var count int64

	countSQL := commonQuery.Session(&gorm.Session{DryRun: true}).Count(&count).Statement.SQL.String()
	fmt.Println("统计SQL:", countSQL)

	sqlStr := commonQuery.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Find(&categoryList)
	})
	fmt.Println("执行sql", sqlStr)

	commonQuery.Limit(int(requestData.Limit)).Offset((int(requestData.Page) - 1) * int(requestData.Limit)).
		Preload("Tools").Find(&categoryList)
	return categoryList, nil

	//var categoryList []model.TCategory
	//q := query.Use(database.DB).TCategory.WithContext(ctx)
}
