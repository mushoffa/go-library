package redis

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/redis.v5"
)

// @Created 30/11/2021
// @Updated
type RedisClientService interface {
	RedisService
	GetInstance() *redis.Client
}

// @Created 16/11/2021
// @Updated 30/11/2021
type RedisClient struct {
	client *redis.Client
}

// @Created 16/11/2021
// @Updated 30/11/2021
func NewRedisClient(host, password string, db int) (RedisClientService, error) {

	if host == "" {
		return nil, fmt.Errorf("Redis host cannot be empty")
	}

	r := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password, // "", empty string if no password set
		DB:       db,       // 0, to use default DB
	})

	if _, err := r.Ping(context.Background()).Result(); err != nil {
		return nil, fmt.Errorf("Failed connecting to redis: %v", err)
	}

	return &RedisClient{r}, nil
}

// @Created 16/11/2021
// @Updated
func (r *RedisClient) GetInstance() *redis.Client {
	return r.client
}

// @Created 16/11/2021
// @Updated
func (r *RedisClient) Decr(key string) *redis.IntCmd {
	return r.client.Decr(context.Background(), key)
}

// @Created 16/11/2021
// @Updated
func (r *RedisClient) DecrBy(key string, decrement int64) *redis.IntCmd {
	return r.client.DecrBy(context.Background(), key, decrement)
}

// @Created 16/11/2021
// @Updated
func (r *RedisClient) Del(keys ...string) *redis.IntCmd {
	return r.client.Del(context.Background(), keys...)
}

// @Created 16/11/2021
// @Updated
func (r *RedisClient) Get(key string) *redis.StringCmd {
	return r.client.Get(context.Background(), key)
}

// @Created 16/11/2021
// @Updated
func (r *RedisClient) Incr(key string) *redis.IntCmd {
	return r.client.Incr(context.Background(), key)
}

// @Created 16/11/2021
// @Updated
func (r *RedisClient) IncrBy(key string, value int64) *redis.IntCmd {
	return r.client.IncrBy(context.Background(), key, value)
}

// @Created 16/11/2021
// @Updated
func (r *RedisClient) Ping() *redis.StatusCmd {
	return r.client.Ping(context.Background())
}

// @Created 16/11/2021
// @Updated
func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.client.Set(context.Background(), key, value, expiration)
}

// @Created 16/11/2021
// @Updated
func (r *RedisClient) Publish(channel string, message interface{}) *redis.IntCmd {
	return r.client.Publish(context.Background(), channel, message)
}

// @Created 16/11/2021
// @Updated
func (r *RedisClient) PSubscribe(keys ...string) *redis.PubSub {
	return r.client.PSubscribe(context.Background(), keys...)
}
