package event

import (
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(name string, ch *amqp.Channel) error {
	switch name {
	case "auth_user", "get_user_by_email", "authenticate_user_session", "get_all_knowledge", "add_users_knowledge",
		"get_percent_knowledge":
		return ch.ExchangeDeclare(
			name,
			"topic",
			true,
			false,
			false,
			false,
			nil,
		)
	default:
		return errors.New("invalid name of channel RabbitMQ")
	}
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
}
