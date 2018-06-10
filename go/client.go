package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func NewClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	return client, err
}

func SetGet(c *redis.Client) {
	err := c.Set("my_key", "Hello World from go client", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := c.Get("my_key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("my_key", val)

	val2, err := c.Get("my_key_2").Result()
	if err == redis.Nil {
		fmt.Println("my_key_2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("my_key_2", val2)
	}
}

func main() {
	client, err := NewClient()
	if err != nil {
		panic(err)
	}
	SetGet(client)
}
