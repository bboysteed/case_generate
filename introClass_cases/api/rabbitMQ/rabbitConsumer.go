package rabbitMQ

import (
	"api/utils/log"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:123456@127.0.0.1:5672")
	handleError(err, "连接失败")
	defer conn.Close()
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
	handleError(err2,"申明通道失败")

	msgs,recErr := ch.Consume(
		quene.Name,
		"consumer1",
		false,
		false,
		false,
		false,
		nil,
		)
	handleError(recErr,"接受消息错误")
	for msg := range msgs {
		log.Info().Msgf("rec message from quene:%s,%s",quene.Name,msg.Body)
		msg.Ack(false)

	}

}
func handleError(err error, message string) {
	if err != nil {
		log.Info().Msgf("%s,err is:%s ", message, err.Error())
	}
}

