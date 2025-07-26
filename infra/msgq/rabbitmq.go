package msgq

import (
	"fmt"
	"gobackend/core/configuration"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQ(cfg configuration.Config) (*amqp.Connection, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d%s",
		cfg.RabbitMQ.User, cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host, cfg.RabbitMQ.Port, cfg.RabbitMQ.VHost)

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
