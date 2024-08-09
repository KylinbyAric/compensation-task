package redis

import (
	"github.com/go-redis/redis/v8"
	"sync"
)

var Client *redis.Client

func GetRedisConnection() *redis.Client {
	return Client
}

func InitRedis() {
	once := sync.Once{}
	once.Do(func() {
		Client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", // Redis地址  todo 抽成配置文件
			Password: "",               // Redis密码，如果没有则为空字符串
			DB:       0,                // 使用默认DB
		})
	})
}
