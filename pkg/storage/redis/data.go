package redis 

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func Read(client *redis.Client, key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
    
	return val, err
}

func Write(client *redis.Client, key, value string) error {
	err := client.Set(ctx, key, value, 0).Err()

	return err
}