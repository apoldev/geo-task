package cache

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(host, port string) *redis.Client {
	// реализуйте создание клиента для Redis

	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
		DB:   0,
	})

}
