package redis 

import (
    "github.com/go-redis/redis/v8"
)

func Client(addr, password string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       db,  // use default DB
    })

	return client
}