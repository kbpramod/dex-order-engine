package main

import (
	"fmt"

	"services/internal/config"
	"services/internal/redis"
	"services/internal/stream/services"
)

func main() {
	config.LoadEnv()
	redis.InitRedis()
	pong, err := redis.Client.Ping(redis.Ctx).Result()

if err != nil {
	panic(err)
}

fmt.Println(pong)
	services.StartStream()
	fmt.Println("🚀 Quote Stream Worker running...")
}