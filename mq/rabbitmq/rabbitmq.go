package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"log"
)

type RabbitMQ struct {
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
}

var MQ *RabbitMQ

func InitRabbitMQ() {
	url := viper.GetString("rabbitmq.url")
	conn, err := amqp091.Dial(url)
	if err != nil {
		log.Println("❌ RabbitMQ链接失败", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		log.Println("❌ RabbitMQ开启通道失败", err)
	}

	MQ = &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}
	log.Println("✅ RabbitMQ链接成功")
}

func Close() {
	if MQ != nil {
		MQ.Channel.Close()
		MQ.Conn.Close()
	}
}
