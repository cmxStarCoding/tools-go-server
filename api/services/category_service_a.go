package services

import (
	"journey/api/validator"
	"journey/common/database"
	"journey/models"
)

// AsService 测试
type AsService struct{}

func (s AsService) GetCategoryList(requestData *validator.GetCategoryListRequest) (map[int]models.CategoryModel, error) {
	db := database.DB
	var categoryList []models.CategoryModel

	db.Limit(int(requestData.Limit)).Offset((int(requestData.Page) - 1) * int(requestData.Limit)).Find(&categoryList)

	categoryMap := make(map[int]models.CategoryModel)

	for _, v2 := range categoryList {
		if int(v2.Pid) == 0 {
			categoryMap[int(v2.ID)] = v2
		}
		// 将子节点加入父节点的 Children 切片中
		if parent, exists := categoryMap[int(v2.Pid)]; exists && int(v2.Pid) > 0 {
			parent.Children = append(parent.Children, v2)
			categoryMap[int(v2.Pid)] = parent
		}
	}

	return categoryMap, nil

}
