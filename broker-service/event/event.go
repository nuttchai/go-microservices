package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // name of the exchange
		"topic",      // topic
		true,         // is durable?
		false,        // auto-delete?
		false,        // internal?
		false,        // no-wait?
		nil,          // argument
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused?
		true,  // exclusive?
		false, // no-wait
		nil,   // argument
	)
}
