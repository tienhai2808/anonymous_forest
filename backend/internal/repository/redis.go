package repository

import (
	"context"
	"time"
)

type RedisRepository interface {
	SetString(ctx context.Context, key string, str string, ttl time.Duration) error

	GetString(ctx context.Context, key string) (string, error)

	IncrementWithTTL(ctx context.Context, key string, ttl time.Duration) (int64, error)

	Decrement(ctx context.Context, key string) error
}
