package implement

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tienhai2808/anonymous_forest/backend/internal/repository"
)

type redisRepositoryImpl struct {
	rdb *redis.Client
}

func NewRedisRepository(rdb *redis.Client) repository.RedisRepository {
	return &redisRepositoryImpl{rdb}
}

func (r *redisRepositoryImpl) SetString(ctx context.Context, key string, str string, ttl time.Duration) error {
	return r.rdb.Set(ctx, key, str, ttl).Err()
}

func (r *redisRepositoryImpl) GetString(ctx context.Context, key string) (string, error) {
	str, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		return "", err
	}

	return str, nil
}

func (r *redisRepositoryImpl) IncrementWithTTL(ctx context.Context, key string, ttl time.Duration) (int64, error) {
	var incrCmd *redis.IntCmd
	if _, err := r.rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		incrCmd = pipe.Incr(ctx, key)
		pipe.ExpireNX(ctx, key, ttl)
		return nil
	}); err != nil {
		return 0, err
	}

	return incrCmd.Val(), nil
}

func (r *redisRepositoryImpl) Decrement(ctx context.Context, key string) error {
	return r.rdb.Decr(ctx, key).Err()
}
