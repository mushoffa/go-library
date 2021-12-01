package redis_test

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
	rr "github.com/mushoffa/go-library/redis"
)

func TestNewRedisClusterClient_Success(t *testing.T) {

	// Change IP & Port of redis cluster according to your setup
	r, err := rr.NewRedisClusterClient(&redis.ClusterOptions{
		Addrs: []string{":6380", ":6381"},
	})
	if err != nil {
		t.Errorf("Error creating redis cluster client: %v", err)
	}
	fmt.Println("Redis PING: ", r.Ping())
}
