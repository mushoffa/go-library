package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/mushoffa/go-library/redis"
)

var (
	redisHost = ":6379"
	redisPassword = ""
	redisDB = 0
)

// @Created 16/11/2021
// @Updated
func TestNewRedisClient_Success(t *testing.T) {

	r, err := redis.NewRedisClient(redisHost, redisPassword, redisDB)
	if err != nil {
		t.Errorf("Error creating redis client: %v", err)
	}

	fmt.Println("Redis PING: ", r.Ping(context.Background()))
}

// @Created 16/11/2021
// @Updated
func TestNewRedisRepository_Success(t *testing.T) {
	repository, err := redis.NewRedisRepository(redisHost, redisPassword, redisDB)
	if err != nil {
		t.Errorf("Error creating redis repository: %v", err)
	}

	fmt.Println("Redis repository PING: ", repository.Ping())
}