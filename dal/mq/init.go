package mq

import (
	"log"

	"github.com/jlu-cow-studio/common/discovery"
	"github.com/segmentio/kafka-go"
)

const (
	Topic_ItemChange  = "item_change"
	Topic_ClientEvent = "client_event"
)

var (
	addressConsumer string
	addressProducer string
)

func Init() {
	log.Println("initialing mq ...")
	addresses, err := discovery.DiscoverMQ()
	if err != nil {
		panic(err)
	}

	addressConsumer = addresses[0]
	addressProducer = addresses[0]

	log.Println("get address ", addressConsumer, addressProducer)

	topics := []string{
		Topic_ItemChange,
		Topic_ClientEvent,
	}

	log.Println("registering topics ...")
	for _, topic := range topics {
		err := RegisterTopic(addressProducer, topic)
		if err != nil {
			panic(err)
		}
	}

	log.Println("mq init done!")
}

func RegisterTopic(address, topic string) error {
	log.Println("creating topic , ", topic)
	conn, err := kafka.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = conn.CreateTopics(topicConfigs...)
	if err != nil {
		return err
	}

	return nil
}
