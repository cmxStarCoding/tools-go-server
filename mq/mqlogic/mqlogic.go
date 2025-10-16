package mqlogic

import (
	"journey/mq/rabbitmq"
	"log"
)

func StartMqTask() {
	// 启动消费者（可以放在单独服务中）
	rabbitmq.Consume("test_queue1", func(msg string) {
		log.Println("👀 Received:", msg)
	})

	rabbitmq.Consume("test_queue2", func(msg string) {
		log.Println("👀 Received11:", msg)
	})

}
