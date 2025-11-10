package main

import (
	"fmt"
	"journey/cmd"
	"journey/common/cache"
	"journey/common/database"
	"journey/common/middleware"
	"journey/common/utils"
	"journey/cron"
	"journey/mq/mqlogic"
	"journey/mq/rabbitmq"
	"journey/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 如果有 CLI 参数，则优先执行 CLI
	if len(os.Args) > 1 {
		cmd.InitCmd()
		return
	}
	var env string
	env = os.Getenv("ENV")
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	// 初始化Gin
	r := gin.Default()
	//加载配置文件
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	r.MaxMultipartMemory = 2 << 20 // 8 MiB
	//设置跨域中间件
	r.Use(middleware.CORSMiddleware())
	//初始化gin日志文件，如果不设置，只会输出在控制台
	utils.SetupLogger()
	// 初始化数据库连接
	database.InitDB()
	//分钟级定时器
	cron.RegisterCron()
	//秒级定时器
	cron.RegisterSecondCron()

	// 初始化redis链接
	cache.InitClient()
	//初始化rabbitmq链接
	rabbitmq.InitRabbitMQ()
	//开启消息队列任务
	mqlogic.StartMqTask()
	// 设置API路由
	routes.SetupRoutes(r)
	//静态资源配置
	r.Static("/static", "../static")
	// 启动服务
	err := r.Run(":8083")
	if err != nil {
		fmt.Println("❌服务启动失败", err)
		return
	}

}
