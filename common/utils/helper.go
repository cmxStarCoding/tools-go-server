package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
	"unicode"
)

var logLock sync.Mutex

func SetupLogger() {
	//设置日志文件
	gin.DisableConsoleColor()
	// 记录到文件。
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func Md5Hash(input string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(input))
	return hex.EncodeToString(md5Hash.Sum(nil))
}

// 生成唯一的随机字符串
func GenerateUniqueRandomString() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomString := make([]byte, 10) // 10个字符长度，你可以根据需要调整
	rand.Seed(time.Now().UnixNano())
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	// 使用时间戳生成唯一性
	uniquePart := time.Now().UnixNano()
	return fmt.Sprintf("%s%d", string(randomString), uniquePart)
}

func StringNumericToUnit(stringNumeric string) uint {
	id64, _ := strconv.ParseUint(stringNumeric, 10, 64)
	return uint(id64)
}

func ContainValue(slice []uint64, value uint64) bool {

	for _, v := range slice {

		if v == value {
			return true
		}
	}
	return false
}

func ToSnakeCase(s string) string {
	var res string
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			res += "_"
		}
		res += string(unicode.ToLower(r))
	}
	return res
}

func WriteLog(filePath, message string) error {
	dateDir := time.Now().Format("20060102")
	logDir := filepath.Join("log", dateDir)

	// 自动创建目录（若不存在）
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return fmt.Errorf("创建日志目录失败: %w", err)
	}

	// 日志文件路径
	fullPath := filepath.Join(logDir, filePath+".log")

	// 加锁防止并发写冲突
	logLock.Lock()
	defer logLock.Unlock()

	// 打开或创建文件
	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("无法打开日志文件: %w", err)
	}
	defer file.Close()

	// 创建日志记录器（标准格式）
	logger := log.New(file, "", log.LstdFlags)
	logMessage := fmt.Sprintf("[%s] %s", filePath, message)
	logger.Println(logMessage)

	return nil
}
