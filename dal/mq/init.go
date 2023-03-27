package mq

import "github.com/jlu-cow-studio/common/discovery"

const (
	Topic_ItemChange  = "item_change"
	Topic_ClientEvent = "client_event"
)

var (
	addressConsumer string
	addressProducer string
)

func Init() {
	addresses, err := discovery.DiscoverMQ()
	if err != nil {
		panic(err)
	}

	addressConsumer = addresses[0]
	addressProducer = addresses[0]
}
