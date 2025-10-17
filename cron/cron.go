package cron

import (
	"github.com/robfig/cron/v3"
	"log"
)

// RegisterCron 分钟级别的定时任务，最小到分钟级别
func RegisterCron() {
	// 创建 cron 实例
	c := cron.New()
	//c := cron.New(cron.WithSeconds()) //秒级别的定时器，只能运行秒级的
	// 每分钟运行一次
	_, _ = c.AddFunc("* * * * *", func() {
		//fmt.Println("分钟级定时器执行了")

		// 在这里执行你的脚本或任务
		//if utils.IsProd() {
		//services.UserService{}.CheckNewUserGiftExpire()
		//}
	})
	// 每五分钟执行一次
	//_, _ = c.AddFunc("*/5 * * * *", func() {
	// 在这里执行你的脚本或任务
	//services.SnapshotService{}.CronSnapTask()
	//})

	// 每五秒执行一次
	//_, _ = c.AddFunc("*/10 * * * * *", func() {
	// 在这里执行你的脚本或任务
	//runYourScript1()
	//services.SnapshotService{}.CronSnapTask()
	//})

	// 输出日志，确保 cron 任务被注册
	//fmt.Println("Cron tasks registered")
	// 启动 cron
	c.Start()
	log.Println("✅ 分钟级定时器启动成功")

}

// RegisterSecondCron 秒级别的定时任务，最小到秒级
func RegisterSecondCron() {
	s := cron.New(cron.WithSeconds()) //秒级别的定时器，只能运行秒级的
	_, _ = s.AddFunc("*/3 * * * * *", func() {
		//fmt.Println("秒级定时器执行了")

	})
	s.Start()

	log.Println("✅ 秒级定时器启动成功")

}
