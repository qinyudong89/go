package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

func initClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "42.194.169.190:26379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func main() {
	rdb := initClient()

	val := rdb.Get(ctx, "db")
	fmt.Printf("val: %v\n", val)
}
