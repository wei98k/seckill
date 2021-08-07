package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()


func main()  {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.2.27:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "test-key-2", 5, 0).Err()
	if err != nil {
		panic(err)
	}

	// res, err := rdb.Exists(ctx, "test-key-99").Result()

	// str := rdb.Exists(ctx, "test-key-99").Val()

	for i:=0;i<10;i++ {
		val := rdb.Decr(ctx, "test-key-2").Val()

		fmt.Println(val)
	}

	fmt.Println("redis run")
}