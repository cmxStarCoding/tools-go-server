package controllers

import (
	"github.com/gin-gonic/gin"
	"journey/api/services"
	"journey/api/validator"
)

type CategoryController struct{}

// GetCategoryList 获取分类列表
func (c *CategoryController) GetCategoryList(ctx *gin.Context) {
	HandleRequest(ctx,
		validator.ValidateGetCategoryList,
		services.CategoryService{}.GetCategoryList,
	)
}

// GetCategoryToolsList 获取分类工具列表
func (c *CategoryController) GetCategoryToolsList(ctx *gin.Context) {
	HandleRequest(ctx,
		validator.ValidateGetCategoryToolsRequest,
		services.CategoryService{}.GetCategoryToolsList,
	)
}
