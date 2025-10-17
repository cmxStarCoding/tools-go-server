package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleRequest 通用函数
func HandleRequest[Req any, Res any](
	ctx *gin.Context,
	validate func(*gin.Context) (Req, error),
	handle func(Req) (Res, error),
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
