package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleRequest 通用函数
func HandleRequest[Req any, Res any](
	ctx *gin.Context,
	validate func(*gin.Context) (Req, error), //Req由validate函数的返回类型来推导
	handle func(*gin.Context, Req) (Res, error), //Res由handle函数的返回类型来推导
) {
	req, err := validate(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := handle(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
