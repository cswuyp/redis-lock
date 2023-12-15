package redis_lock

import (
	"context"
	"time"
)

func (s *redisLockClient) Lock(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	ret, err := s.client.SetNX(ctx, key, value, expiration).Result()
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func (s *redisLockClient) LockWithOpt(ctx context.Context, key string, value interface{}, expiration time.Duration, watchDog bool) (bool, error) {
	s.Lock()
}

// todo 看门狗
func (s *redisLockClient) watchDog(ctx context.Context)
