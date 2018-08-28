package redis

import (
	"github.com/go-redis/redis"
	"strings"
)

func NewRedisClusterConfig(hosts string) *redis.ClusterOptions {
	addr := strings.Split(hosts, ",")
	return &redis.ClusterOptions{
		Addrs:      addr,
		MaxRetries: 2,
	}
}

func NewRedisClusterClient(opts *redis.ClusterOptions) *redis.ClusterClient {
	if opts == nil {
		panic("Client options can not be nill")
	}
	return redis.NewClusterClient(opts)

}
