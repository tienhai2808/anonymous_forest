package repository

import (
	"context"
	"time"
)

type RedisRepository interface {
	SaveString(ctx context.Context, key string, str string, ttl time.Duration) error
}