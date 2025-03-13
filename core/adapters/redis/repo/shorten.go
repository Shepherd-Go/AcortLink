package repo

import (
	"acortlink/core/domain/models"
	"acortlink/core/domain/ports"
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type shortenRepoRedis struct {
	redisClient *redis.Client
}

func NewShortenRepoRedis(redisClient *redis.Client) ports.ShortenRepoRedis {
	return &shortenRepoRedis{redisClient}
}

func (r *shortenRepoRedis) Save(ctx context.Context, key string, url models.URLResponse, time time.Duration) error {

	mUrl, _ := json.Marshal(url)

	if err := r.redisClient.Set(ctx, key, mUrl, time).Err(); err != nil {
		return err
	}

	return nil
}

func (r *shortenRepoRedis) SearchUrl(ctx context.Context, path string) (models.URLResponse, error) {

	var url = models.URLResponse{}

	val, err := r.redisClient.Get(ctx, path).Result()
	if err != nil {
		return models.URLResponse{}, err
	}

	_ = json.Unmarshal([]byte(val), &url)

	return url, nil
}
