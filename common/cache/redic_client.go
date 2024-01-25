package cache

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

var RedisClient *redis.Client

func InitClient() {
	// 连接到 Redis 服务器
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",  // Redis 服务器地址
		Password:     "",                // Redis 服务器密码
		DB:           0,                 // 默认数据库
		PoolSize:     10,                // 连接池大小
		MinIdleConns: 5,                 // 最小空闲连接数
		IdleTimeout:  240 * time.Second, // 空闲连接的超时时间
	})

	// 检查连接是否成功
	pong, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalln("Failed to connect to Redis:", err)
	} else {
		log.Println("Connected to Redis:", pong)
	}
}
