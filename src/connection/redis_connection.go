package connection

import (
	"github.com/redis/go-redis/v9"
)

type RedisConnection struct {
    Client *redis.Client
}

func GetRedisConnection(addr string) *RedisConnection {
    client := redis.NewClient(&redis.Options{
        Addr: addr,
        Password: "",
        DB: 0,
    })
    return &RedisConnection{client}
}
