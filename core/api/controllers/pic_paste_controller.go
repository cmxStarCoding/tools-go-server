package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tools/core/api/services"
	"tools/core/api/validator/pic"
)

type PicPasteController struct{}

// PicPaste 图片贴图能力
func (c PicPasteController) PicPaste(ctx *gin.Context) {
	request, requestErr := pic.ValidateRequest(ctx)
	if requestErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": requestErr.Error()})
		return
	}
	//
	var returnData map[string]string
	go services.PicPasteService{}.DoTask(request)

	ctx.JSON(http.StatusBadRequest, returnData)
	return
}
