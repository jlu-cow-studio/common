package discovery

import (
	"errors"
	"fmt"
)

const redisServiceName = "cowstudio/redis"
const mqServiceName = "cowstudio/mq"

func Discover(serviceName, tag string) ([]string, error) {

	services, _, err := consulClient.Catalog().Service(serviceName, tag, nil)
	if err != nil {
		return nil, nil
	}
	serviceAddress := []string{}
	for _, service := range services {
		serviceAddress = append(serviceAddress, fmt.Sprintf("%v:%v", service.ServiceAddress, service.ServicePort))
	}
	if len(serviceAddress) == 0 {
		return nil, errors.New("no service found")
	}

	return serviceAddress, nil
}

func DiscoverRedis() ([]string, error) {
	return Discover(redisServiceName, "")
}

func DiscoverMQ() ([]string, error) {
	return Discover(mqServiceName, "")
}
