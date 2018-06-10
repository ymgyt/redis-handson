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

func UpVote(c *redis.Client, id string) {
	var key = "article:" + id + ":votes"
	c.Incr(key)
}

func DownVote(c *redis.Client, id string) {
	var key = "article:" + id + ":votes"
	c.Decr(key)
}

func ShowResult(c *redis.Client, id string) {
	var headlineKey = "article:" + id + "headline"
	var voteKey = "article:" + id + ":votes"
	v, err := c.MGet(headlineKey, voteKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}

func main() {
	client, err := NewClient()
	if err != nil {
		panic(err)
	}

	UpVote(client, "12345")
	UpVote(client, "12345")
	UpVote(client, "12345")
	UpVote(client, "12345")
	UpVote(client, "12345")
	UpVote(client, "12345")

	ShowResult(client, "12345")
}
