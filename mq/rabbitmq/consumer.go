// pkg/rabbitmq/consumer.go
package rabbitmq

import (
	"log"
)

func Consume(queueName string, handler func(msg string)) error {
	ch := MQ.Channel

	q, err := ch.QueueDeclare(
		queueName, true, false, false, false, nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,  // 自动ACK
		false, // 非独占
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			handler(string(d.Body))
		}
	}()
	log.Printf("👂 Listening on queue [%s]", queueName)
	return nil
}
