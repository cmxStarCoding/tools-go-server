package services

import (
	"tools/common/database"
	"tools/core/api/models"
	"tools/core/api/validator/tools"
)

type ToolsService struct{}

func (s ToolsService) GetToolsList(requestData *tools.GetToolListRequest) ([]models.ToolsModel, error) {
	var toolsMap []models.ToolsModel
	query := database.DB.Limit(int(requestData.Limit)).Offset((int(requestData.Page) - 1) * int(requestData.Limit))
	if requestData.CategoryId != 0 {
		query = query.Where("category_id = ?", requestData.CategoryId)
	}

	if requestData.IsRecommend == 1 {
		query = query.Where("is_recommend = ?", 1)
	}

	query.Find(&toolsMap)
	return toolsMap, nil
}
