package infra

import "github.com/redis/go-redis/v9"


func NewRedisClient(url string) *redis.Client {
    return redis.NewClient(&redis.Options{
        Addr: url,
    })
}