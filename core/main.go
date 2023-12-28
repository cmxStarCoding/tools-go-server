package main

import (
	"github.com/gin-gonic/gin"
	"tools/common/database"
	"tools/common/utils"
	"tools/core/api/v1"
)

func main() {
	// 初始化Gin
	r := gin.Default()
	//初始化日志文件
	utils.SetupLogger()
	// 初始化数据库连接
	database.InitDB()
	// 设置API路由
	v1.SetupRoutes(r)
	// 启动服务
	r.Run(":8080")
}