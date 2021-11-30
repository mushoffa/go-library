package redis

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/redis.v5"
)

// @Created 30/11/2021
// @Updated
type RedisClusterService interface {
	RedisService
	GetInstance() *redis.ClusterClient
}

// @Created 30/11/2021
// @Updated
type RedisCluster struct {
	client *redis.ClusterClient
}

// @Created 16/11/2021
// @Updated 30/11/2021
func NewRedisClusterClient(opt *redis.ClusterOptions) (RedisClusterService, error) {
	r := redis.NewClusterClient(opt)

	if _, err := r.Ping(context.Background()).Result(); err != nil {
		return nil, fmt.Errorf("Failed connecting to redis cluser: %v", err)
	}

	return &RedisCluster{r}, nil
}

// @Created 30/11/2021
// @Updated
func (r *RedisCluster) GetInstance() *redis.ClusterClient {
	return r.client
}

// @Created 30/11/2021
// @Updated
func (r *RedisCluster) Decr(key string) *redis.IntCmd {
	return r.client.Decr(context.Background(), key)
}

// @Created 30/11/2021
// @Updated
func (r *RedisCluster) DecrBy(key string, decrement int64) *redis.IntCmd {
	return r.client.DecrBy(context.Background(), key, decrement)
}

// @Created 30/11/2021
// @Updated
func (r *RedisCluster) Del(keys ...string) *redis.IntCmd {
	return r.client.Del(context.Background(), keys...)
}

// @Created 30/11/2021
// @Updated
func (r *RedisCluster) Get(key string) *redis.StringCmd {
	return r.client.Get(context.Background(), key)
}

// @Created 30/11/2021
// @Updated
func (r *RedisCluster) Incr(key string) *redis.IntCmd {
	return r.client.Incr(context.Background(), key)
}

// @Created 30/11/2021
// @Updated
func (r *RedisCluster) IncrBy(key string, value int64) *redis.IntCmd {
	return r.client.IncrBy(context.Background(), key, value)
}

// @Created 30/11/2021
// @Updated
func (r *RedisCluster) Ping() *redis.StatusCmd {
	return r.client.Ping(context.Background())
}

// @Created 30/11/2021
// @Updated
func (r *RedisCluster) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.client.Set(context.Background(), key, value, expiration)
}

// @Created 30/11/2021
// @Updated
func (r *RedisCluster) Publish(channel string, message interface{}) *redis.IntCmd {
	return r.client.Publish(context.Background(), channel, message)
}

// @Created 30/11/2021
// @Updated
func (r *RedisCluster) PSubscribe(keys ...string) *redis.PubSub {
	return r.client.PSubscribe(context.Background(), keys...)
}
