package controllers

import (
	"journey/api/services"
	"journey/api/validator"
	"journey/models"

	"github.com/gin-gonic/gin"
)

type CategoryController struct{}

// GetCategoryList 获取分类列表
func (c *CategoryController) GetCategoryList(ctx *gin.Context) {
	HandleRequest(ctx,
		validator.ValidateGetCategoryList,
		func(ctx *gin.Context, req *validator.GetCategoryListRequest) ([]models.CategoryModel, error) {
			return services.CategoryService{}.GetCategoryList(req)
		},
	)
}

// GetCategoryToolsList 获取分类工具列表
func (c *CategoryController) GetCategoryToolsList(ctx *gin.Context) {
	HandleRequest(ctx,
		validator.ValidateGetCategoryToolsRequest,
		func(ctx *gin.Context, req *validator.GetCategoryToolsRequest) ([]models.CategoryModel, error) {
			return services.CategoryService{}.GetCategoryToolsList(req)
		},
	)
}
