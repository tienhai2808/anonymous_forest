package implement

import (
	"context"
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

func (r *redisRepositoryImpl) SaveString(ctx context.Context, key string, str string, ttl time.Duration) error {
	return r.rdb.Set(ctx, key, str, ttl).Err()
}