package controllers

import (
	"github.com/gin-gonic/gin"
	"journey/api/services"
)

type WebsocketController struct{}

func (s WebsocketController) MyWs(ctx *gin.Context) {
	go services.WebsocketService{}.Run()
	// 调用用户服务获取用户信息
	services.WebsocketService{}.MyWs(ctx.Writer, ctx.Request)
}
