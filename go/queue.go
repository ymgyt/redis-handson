package main

import (
	"fmt"
	"time"

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

type Queue struct {
	name    string
	client  *redis.Client
	key     string
	timeout time.Duration
}

func NewQueue(c *redis.Client, name string) *Queue {
	return &Queue{
		name:    name,
		client:  c,
		key:     "queues:" + name,
		timeout: time.Duration(0),
	}
}

func (q *Queue) Size() int64 {
	return q.client.LLen(q.key).Val()
}

func (q *Queue) Push(data interface{}) {
	q.client.LPush(q.key, data)
}

func (q *Queue) Pop() interface{} {
	return q.client.BRPop(q.timeout, q.key).Val()
}

func main() {
	c, err := NewClient()
	if err != nil {
		panic(err)
	}
	q := NewQueue(c, "queue")

	n := 5
	for i := 0; i < n; i++ {
		q.Push(fmt.Sprintf("Hello %d", i))
	}

	for i := 0; i < n; i++ {
		fmt.Println(q.Pop())
	}

	c.Close()
}
