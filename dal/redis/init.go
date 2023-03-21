package redis

import (
	"github.com/go-redis/redis"
	"github.com/jlu-cow-studio/common/discovery"
)

const RedisPassword = "cowstudio"

var DB *redis.Client

func Init() {

	redisAddrs, err := discovery.DiscoverRedis()
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddrs[0],
		Password: RedisPassword,
		DB:       0, // use default DB
	})

	DB = rdb
}
