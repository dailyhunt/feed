package feed_redis

import (
	"fmt"
	"testing"
)

func TestRedisConn(t *testing.T) {

	//callRedis()

}

func callRedis() {
	hosts := "192.168.25.12:7001"
	conf := NewRedisClusterConfig(hosts)
	client := NewRedisClusterClient(conf)
	keys := client.HKeys("analytics.meta.news_item_group.N7008624")
	fmt.Println(keys.Val())
}
