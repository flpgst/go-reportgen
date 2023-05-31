package interfaces

import "github.com/streadway/amqp"

type RabbitMQInterface interface {
	OpenChannel() *amqp.Channel
	Consume(ch *amqp.Channel, out chan<- amqp.Delivery, queue string) error
}
