package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleRequest 通用函数
func HandleRequest[Req any, Res any](
	ctx *gin.Context,
	validate func(*gin.Context) (Req, error), //Req由validate函数的返回类型来推导
	handle func(Req) (Res, error), //Res由handle函数的返回类型来推导
) {
	req, err := validate(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := handle(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
