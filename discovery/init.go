package discovery

import "github.com/hashicorp/consul/api"

const CONSUL_ADDRESS = "cowstudio.wayne-lee.cn:3001"

var consulClient *api.Client

func Init() {
	config := api.DefaultConfig()
	config.Address = CONSUL_ADDRESS
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	consulClient = client
}
