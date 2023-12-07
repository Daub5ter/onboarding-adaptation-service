package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	if err := ch.ExchangeDeclare("auth_user", "topic", true, false, false, false, nil); err != nil {
		return err
	}
	if err := ch.ExchangeDeclare("get_user_by_email", "topic", true, false, false, false, nil); err != nil {
		return err
	}
	if err := ch.ExchangeDeclare("authenticate_user_session", "topic", true, false, false, false, nil); err != nil {
		return err
	}
	if err := ch.ExchangeDeclare("get_all_knowledge", "topic", true, false, false, false, nil); err != nil {
		return err
	}
	if err := ch.ExchangeDeclare("add_users_knowledge", "topic", true, false, false, false, nil); err != nil {
		return err
	}
	if err := ch.ExchangeDeclare("get_percent_knowledge", "topic", true, false, false, false, nil); err != nil {
		return err
	}

	return nil
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
