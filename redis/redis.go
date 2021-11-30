package redis

import (
	"time"

	"github.com/go-redis/redis/v8"
)

// @Created 16/11/2021
// @Updated 30/11/2021
type RedisService interface {
	Decr(string) *redis.IntCmd

	DecrBy(string, int64) *redis.IntCmd

	// Context, Keys
	Del(...string) *redis.IntCmd

	// Context, Key
	Get(string) *redis.StringCmd

	Incr(string) *redis.IntCmd

	IncrBy(string, int64) *redis.IntCmd

	Ping() *redis.StatusCmd

	// Context, Keys, Value, Expiration
	Set(string, interface{}, time.Duration) *redis.StatusCmd

	RedisPubSub
}

// @Created 16/11/2021
// @Updated
type RedisPubSub interface {
	Publish(string, interface{}) *redis.IntCmd
	PSubscribe(...string) *redis.PubSub
}

// @Created 16/11/2021
// @Updated
// func NewRedisRepository(host, password string, db int) (RedisService, error) {
// 	r, err := NewRedisClient(host, password, db)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &RedisClient{r}, nil
// }
