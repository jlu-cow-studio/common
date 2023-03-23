package rpc

import (
	"sync"

	"github.com/jlu-cow-studio/common/discovery"
	"google.golang.org/grpc"
)

var connPool = map[string]*grpc.ClientConn{}
var mutex = &sync.Mutex{}

func GetConn(serviceName string) (*grpc.ClientConn, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if conn, ok := connPool[serviceName]; ok {
		return conn, nil
	}

	addrs, err := discovery.Discover(serviceName, "")
	if err != nil {
		return nil, err
	}

	addr := addrs[0]

	if conn, err := grpc.Dial(addr, grpc.WithInsecure()); err != nil {
		return nil, err
	} else {
		connPool[serviceName] = conn
		return conn, nil
	}
}
