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
	var intCmd *redis.IntCmd
	if _, err := r.rdb.TxPipelined(ctx, func(p redis.Pipeliner) error {
		intCmd = p.Incr(ctx, key)
		p.ExpireNX(ctx, key, ttl)
		return nil
	}); err != nil {
		return 0, err
	}

	return intCmd.Val(), nil
}

func (r *redisRepositoryImpl) Decrement(ctx context.Context, key string) error {
	return r.rdb.Decr(ctx, key).Err()
}

func (r *redisRepositoryImpl) SetAddWithTTL(ctx context.Context, key string, str string, ttl time.Duration) error {
	var intCmd *redis.IntCmd
	if _, err := r.rdb.TxPipelined(ctx, func(p redis.Pipeliner) error {
		intCmd = p.SAdd(ctx, key, str)
		p.ExpireNX(ctx, key, ttl)
		return nil
	}); err != nil {
		return err
	}

	return intCmd.Err()
}

func (r *redisRepositoryImpl) SetMembers(ctx context.Context, key string) ([]string, error) {
	return r.rdb.SMembers(ctx, key).Result()
}
