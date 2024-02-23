package database

import "github.com/go-redis/redis/v8"

func RedisConnection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis_db:6379",
		Password: "",
		DB:       0,
	})

	return rdb
}

// EOF
