package initialization

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tienhai2808/anonymous_forest/backend/config"
)

func InitCache(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:      cfg.Cache.CAddr,
		DB:        cfg.Cache.CDb,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("ping tới Redis thất bại: %v", err)
	}

	return rdb, nil
}
