package amqp

import (
	"github.com/rabbitmq/amqp091-go"
)

type Amqp struct {
	Connect *amqp091.Connection `json:"connect"`
	Dsn     string              `json:"dsn"`
}

func InitAMQP(dsn string) (*Amqp, error) {
	var am Amqp
	am.Dsn = dsn

	conn, err := amqp091.Dial(dsn)
	if err != nil {
		return &am, err
	}
	am.Connect = conn
	return &am, nil
}

func (amqp *Amqp) Close() error {
	err := amqp.Connect.Close()
	if err != nil {
		return err
	}
	return nil
}
