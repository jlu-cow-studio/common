package mq

import (
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
