package queue

import (
	"fmt"
	"log"
	"reflect"
)

const (
	RabbitMQ QueueType = iota
)

func New(qt QueueType, cfg any) (q *Queue, err error) {
	rt := reflect.TypeOf(cfg)

	switch qt {
	case RabbitMQ:
		if rt.Name() != "RabbitMQConfig" {
			return nil, fmt.Errorf("config is not compatible with queue time: %s", rt.Name())
		}
		conn, err := newRabbitConn(cfg.(RabbitMQConfig))
		if err != nil {
			return nil, err
		}
		q.qc = conn
	default:
		log.Fatalf("Unsupported queue type: %v", qt)
	}

	return
}

type QueueType int

type QueueConnection interface {
	Publish([]byte) error
	Consume(chan<- QueueDTO) error
}

type Queue struct {
	qc QueueConnection
}

func (q *Queue) Publish(msg []byte) error {
	return q.qc.Publish(msg)
}

func (q *Queue) Consume(cdto chan<- QueueDTO) error {
	return q.qc.Consume(cdto)
}
