package repo

import (
	"acortlink/core/domain/ports"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type shortenRepoRedis struct {
	redisClient *redis.Client
}

func NewShortenRepoRedis(redisClient *redis.Client) ports.ShortenRepoRedis {
	return &shortenRepoRedis{redisClient}
}

func (r *shortenRepoRedis) CreateShorten(ctx context.Context, key string, value interface{}, time time.Duration) error {

	if err := r.redisClient.Set(ctx, key, value, time).Err(); err != nil {
		return err
	}

	return nil
}

func (r *shortenRepoRedis) SearchUrl(ctx context.Context, path string) (string, error) {

	val, err := r.redisClient.Get(ctx, path).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}
