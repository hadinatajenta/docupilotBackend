package cache

import (
	"fmt"
	"gobackend/core/configuration"

	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg configuration.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
}
