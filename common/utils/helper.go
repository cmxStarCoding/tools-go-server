package utils

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func SetupLogger()  {
	//设置日志文件
	gin.DisableConsoleColor()
	// 记录到文件。
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
