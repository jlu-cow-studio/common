package rpc

import (
	"sync"

	"github.com/jlu-cow-studio/common/dal/rpc/feed_service"
	"github.com/jlu-cow-studio/common/dal/rpc/pack"
	"github.com/jlu-cow-studio/common/dal/rpc/product_core"
	"github.com/jlu-cow-studio/common/dal/rpc/tag_core"
	"github.com/jlu-cow-studio/common/dal/rpc/trade_core"
	"github.com/jlu-cow-studio/common/dal/rpc/user_core"
	"github.com/jlu-cow-studio/common/discovery"
	"google.golang.org/grpc"
)

var connPool = map[string]*grpc.ClientConn{}
var mutex = &sync.Mutex{}

const (
	ServiceName_UserCore        = "cowstudio/user-core"
	ServiceName_ProductCore     = "cowstudio/product-core"
	ServiceName_TagCore         = "cowstudio/tag-core"
	ServiceName_TradeCore       = "cowstudio/trade-core"
	ServiceName_DataCollector   = "cowstudio/Data-collector"
	ServiceName_Pack            = "cowstudio/pack"
	ServiceName_Feed            = "cowstudio/feed"
	ServiceName_RecommenderCore = "cowstudio/recommender-core"
)

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

func GetUserCoreCli() (user_core.UserCoreServiceClient, error) {
	conn, err := GetConn(ServiceName_UserCore)
	if err != nil {
		return nil, err
	}

	return user_core.NewUserCoreServiceClient(conn), nil

}

func GetProductCoreCli() (product_core.ProductCoreServiceClient, error) {
	conn, err := GetConn(ServiceName_ProductCore)
	if err != nil {
		return nil, err
	}

	return product_core.NewProductCoreServiceClient(conn), nil
}

func GetTagCoreCli() (tag_core.TagCoreServiceClient, error) {
	conn, err := GetConn(ServiceName_TagCore)
	if err != nil {
		return nil, err
	}

	return tag_core.NewTagCoreServiceClient(conn), nil
}

func GetTradeCoreCli() (trade_core.TradeCoreServiceClient, error) {
	conn, err := GetConn(ServiceName_TradeCore)
	if err != nil {
		return nil, err
	}

	return trade_core.NewTradeCoreServiceClient(conn), nil
}

// func GetDataCollectorCli() (data_collector.DataCollectorServiceClient, error) {
// 	conn, err := GetConn(ServiceName_DataCollector)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data_collector.NewDataCollectorServiceClient(conn), nil
// }

func GetPackCli() (pack.PackServiceClient, error) {
	conn, err := GetConn(ServiceName_Pack)
	if err != nil {
		return nil, err
	}

	return pack.NewPackServiceClient(conn), nil
}

func GetFeedCli() (feed_service.FeedServiceClient, error) {
	conn, err := GetConn(ServiceName_Feed)
	if err != nil {
		return nil, err
	}

	return feed_service.NewFeedServiceClient(conn), nil
}

// func GetRecommenderCoreCli() (recommender_core.RecommenderCoreServiceClient, error) {
// 	conn, err := GetConn(ServiceName_RecommenderCore)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return recommender_core.NewRecommenderCoreServiceClient(conn), nil
// }
