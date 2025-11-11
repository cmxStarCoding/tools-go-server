package cron

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

// RunWithLock 执行任务前尝试加锁，防止任务重叠
func RunWithLock(rdb *redis.Client, key string, expire time.Duration, job func()) {
	ok, err := rdb.SetNX(key, "locked", expire).Result()
	if err != nil {
		fmt.Println("Redis error:", err)
		return
	}
	if !ok {
		fmt.Println("上一次任务未完成，跳过执行:", key)
		return
	}

	defer rdb.Del(key) // 执行完释放锁
	job()
}
