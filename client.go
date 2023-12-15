package redis_lock

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisLockClient struct {
	client *redis.Client
}

func NewRedisLocker(client *redis.Client) RedisLock {
	return &redisLockClient{
		client: client,
	}
}

type RedisLock interface {
	Lock(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
}
