package mq

import (
	"context"
	"io"
	"log"

	"github.com/jlu-cow-studio/common/discovery"
	"github.com/segmentio/kafka-go"
)

func Listen(ctx context.Context, topic string, f func(kafka.Message) error) error {

	addresses, err := discovery.DiscoverMQ()
	if err != nil {
		return err
	}

	address := addresses[0]

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{address},
		Topic:   topic,
	})

	for {
		if msg, err := r.ReadMessage(ctx); err != nil {
			if err == io.EOF {
				log.Println("conn closed!")
				return err
			}
			log.Println("error when getting message, reload.. ,", err.Error())
			r.SetOffset(msg.Offset)
		} else if errf := f(msg); errf != nil {
			log.Println("error when calling func, reload.... , ", errf.Error())
			r.SetOffset(msg.Offset)
		}
	}
}
