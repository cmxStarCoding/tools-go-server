package controllers

import (
	"github.com/gin-gonic/gin"
	"journey/api/services"
	"journey/api/validator"
	"net/http"
)

// CategoryController 分类控制器
type CategoryController struct{}

// GetCategoryList 获取分类列表
func (c CategoryController) GetCategoryList(ctx *gin.Context) {
	request, err := validator.ValidateGetCategoryList(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list, listErr := services.CategoryService{}.GetCategoryList(request)
	if listErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": listErr.Error()})
		return
	}
	// 返回JSON数据
	ctx.JSON(200, list)
}

// GetCategoryToolsList 获取分类工具列表
func (c CategoryController) GetCategoryToolsList(ctx *gin.Context) {
	request, err := validator.ValidateGetCategoryToolsRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list, listErr := services.CategoryService{}.GetCategoryToolsList(request)
	if listErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": listErr.Error()})
		return
	}
	// 返回JSON数据
	ctx.JSON(200, list)
}
