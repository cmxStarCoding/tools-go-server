// pkg/rabbitmq/producer.go
package rabbitmq

import (
	"context"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

var ExchangeName = "journey_exchange"

func Publish(queueName string, body string) error {
	ch := MQ.Channel

	// 声明队列（幂等操作）
	_, err := ch.QueueDeclare(
		queueName, // 队列名
		true,      // 是否持久化
		false,     // 是否自动删除
		false,     // 是否独占
		false,     // 是否阻塞
		nil,       // 额外参数
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		ExchangeName, // exchange
		queueName,    // routing key
		false,        // mandatory
		false,        // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}

	log.Printf("📨 Sent message to queue [%s]: %s", queueName, body)
	return nil
}
