package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQConn struct {
	user     string
	password string
	host     string
	port     string
}

func NewRabbitMQConn(user, password, host, port string) *RabbitMQConn {
	return &RabbitMQConn{
		user: user, password: password, host: host, port: port,
	}
}

func (r *RabbitMQConn) OpenChannel() *amqp.Channel {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", r.user, r.password, r.host, r.port))
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch
}

func (r *RabbitMQConn) Consume(ch *amqp.Channel, out chan<- amqp.Delivery, queue string) error {
	msgs, err := ch.Consume(
		queue,
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for msg := range msgs {
		out <- msg
	}
	return nil
}
