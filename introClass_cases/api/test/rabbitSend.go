package main

import (
	"api/utils/log"
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:123456@127.0.0.1:5672")
	handleError(err, "连接失败")
	ch, err1 := conn.Channel()
	handleError(err1, "新建通道失败")
	quene, err2 := ch.QueueDeclare(
		"testquene",
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err2, "申明通道失败")
	for {
		msgBody := ""
		fmt.Scanln(&msgBody)
		err = ch.Publish(
			"",
			quene.Name,
			false,
			false,
			amqp.Publishing{
				ContentType:  "text/plain",
				DeliveryMode: amqp.Persistent,
				Body:         []byte(msgBody),
			},
		)
	}

}

func handleError(err error, message string) {
	if err != nil {
		log.Info().Msgf("%s,err is:%s ", message, err.Error())
	}
}
