package main

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/valkey-io/valkey-go"
)

func main() {
}

func NewRedisClient(ctx context.Context) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := cli.Ping(ctx).Result()
	if err != nil {
		panic("redis client connection error")
	}
	return cli
}

func NewValkeyClient() valkey.Client {
	cli, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{"localhost:16379"},
	})
	if err != nil {
		panic("valkey client connection error")
	}
	return cli
}
