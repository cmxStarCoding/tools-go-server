package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"tools/common/cache"
	"tools/common/config"
	"tools/common/database"
	"tools/common/middleware"
	"tools/common/utils"
	"tools/core/api/v1"
)

func main() {
	var env string
	flag.StringVar(&env, "env", "local", "设置环境")
	// 解析启动的命令行参数
	flag.Parse()
	// 初始化Gin
	r := gin.Default()
	//加载配置文件
	projectConfig := config.InitConfig(env)
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	r.MaxMultipartMemory = 2 << 20 // 8 MiB

	r.Use(middleware.CORSMiddleware())
	//初始化日志文件
	utils.SetupLogger()
	// 初始化数据库连接
	database.InitDB(projectConfig)
	// 初始化redis链接
	cache.InitClient(projectConfig)
	// 设置API路由
	v1.SetupRoutes(r)
	//静态资源配置
	r.Static("/static", "../static")
	// 启动服务
	r.Run(":8083")
}
