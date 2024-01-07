package redis

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"product-service/pkg/config"
)

func NewClient(conf *config.Config, log *zap.Logger) *redis.Client {
	log.Info("Connecting to redis...")
	return redis.NewClient(&redis.Options{
		Addr:     conf.Redis.URI,
		Password: "",
		DB:       0,
	})
}
