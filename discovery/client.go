package discovery

import "fmt"

func Discover(serviceName, tag string) ([]string, error) {

	services, _, err := consulClient.Catalog().Service(serviceName, tag, nil)
	if err != nil {
		return nil, nil
	}
	serviceNames := []string{}
	for _, service := range services {
		serviceNames = append(serviceNames, fmt.Sprintf("%v:%v", service.ServiceAddress, service.ServicePort))
	}

	return serviceNames, nil
}
