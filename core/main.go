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
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	r.MaxMultipartMemory = 2 << 20 // 8 MiB
	//初始化日志文件
	utils.SetupLogger()
	// 初始化数据库连接
	database.InitDB()
	// 设置API路由
	v1.SetupRoutes(r)
	//静态资源配置
	//r.Static("/static", "../static")
	// 启动服务
	r.Run(":8080")
}
