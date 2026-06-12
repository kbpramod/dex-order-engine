package redis

import (
	"context"

	"services/internal/config"

	goredis "github.com/redis/go-redis/v9"
)

var (
	Client *goredis.Client
	Ctx    = context.Background()
)

func InitRedis() {
	Client = goredis.NewClient(&goredis.Options{
		Addr: config.GetEnv("REDIS_HOST") + ":" +
			config.GetEnv("REDIS_PORT"),
		Password: config.GetEnv("REDIS_PASSWORD"),
		DB: 0,
	})
}