package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/segmentio/kafka-go"
)

var writerMap = map[string]*kafka.Writer{}

var lock sync.Mutex

func GetWritter(topic string) *kafka.Writer {
	lock.Lock()
	defer lock.Unlock()

	if writer, ok := writerMap[topic]; ok {
		return writer
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{addressProducer},
		Topic:   topic,
	})

	writerMap[topic] = writer

	return writer
}

func SendMessage(ctx context.Context, topic string, msg interface{}) error {
	writter := GetWritter(topic)
	if writter == nil {
		return fmt.Errorf("unknown topic")
	}
	val, err := json.Marshal(msg)
	if err != nil {
		return nil
	}

	return writter.WriteMessages(ctx, kafka.Message{
		Value: val,
	})
}
