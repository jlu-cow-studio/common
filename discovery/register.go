package discovery

import (
	"github.com/hashicorp/consul/api"
)

const (
	CHECK_TIMEOUT                           = "5s"
	CHECK_INTERVAL                          = "10s"
	CHECK_DEREGISTER_CRITICAL_SERVICE_AFTER = "30s"
)

func Register(serviceName, address, httpCheck, rpcCheck string, port int) error {
	registration := &api.AgentServiceRegistration{}
	registration.Name = serviceName
	registration.Address = address
	registration.Port = port

	if httpCheck != "" {
		check := &api.AgentServiceCheck{}
		check.HTTP = httpCheck
		check.Timeout = CHECK_TIMEOUT
		check.Interval = CHECK_INTERVAL
		check.DeregisterCriticalServiceAfter = CHECK_DEREGISTER_CRITICAL_SERVICE_AFTER
		registration.Check = check
	}

	err := consulClient.Agent().ServiceRegister(registration)
	return err
}
