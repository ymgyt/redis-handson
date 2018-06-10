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

func SaveLink(c *redis.Client, id, author, title, link string) {
	c.HMSet("link:"+id, map[string]interface{}{
		"author": author,
		"title":  title,
		"link":   link,
		"score":  0,
	})
}

func UpVote(c *redis.Client, id string) {
	vote(c, id, 1)
}

func DownVote(c *redis.Client, id string) {
	vote(c, id, -1)
}

func vote(c *redis.Client, id string, n int) {
	c.HIncrBy("link:"+id, "score", int64(n))
}

func ShowDetail(c *redis.Client, id string) {
	fmt.Println(c.HGetAll("link:" + id).Val())
}

func main() {
	c, err := NewClient()
	if err != nil {
		panic(err)
	}

	SaveLink(c, "1", "yuta", "golang", "aaaaa")
	UpVote(c, "1")
	UpVote(c, "1")

	ShowDetail(c, "1")
}
