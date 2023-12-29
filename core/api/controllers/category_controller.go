package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tools/core/api/services"
	"tools/core/api/validator/category"
)

// CategoryController 分类控制器
type CategoryController struct{}

// GetCategoryList 获取分类列表
func (c *CategoryController) GetCategoryList(ctx *gin.Context) {
	request, err := category.ValidateGetCategoryList(ctx)
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
