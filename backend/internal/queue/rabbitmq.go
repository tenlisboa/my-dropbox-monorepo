package queue

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConfig struct {
	URL       string
	TopicName string
	Timeout   time.Time
}

func newRabbitConn(cfg RabbitMQConfig) (rc *RabbitConnection, err error) {
	rc.cfg = cfg
	rc.conn, err = amqp.Dial(rc.cfg.URL)
	return rc, err
}

type RabbitConnection struct {
	cfg  RabbitMQConfig
	conn *amqp.Connection
}

func (rc *RabbitConnection) Publish(msg []byte) error {
	c, err := rc.conn.Channel()
	if err != nil {
		return err
	}
	defer c.Close()

	mp := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "text/plain",
		Body:         msg,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return c.PublishWithContext(ctx, "", rc.cfg.TopicName, false, false, mp)
}

func (rc *RabbitConnection) Consume(cdto chan<- QueueDTO) error {
	ch, err := rc.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rc.cfg.TopicName, // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	for d := range msgs {
		dto := QueueDTO{}
		dto.Unmarshal(d.Body)

		cdto <- dto
	}

	return nil
}
