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

func MarkDealAsSent(c *redis.Client, dealID, userID string) {
	c.SAdd(dealID, userID)
}

func SendDealIfNotSent(c *redis.Client, dealID, userID string) {
	if c.SIsMember(dealID, userID).Val() {
		fmt.Println("Deal", dealID, "was already send to user", userID)
	} else {
		fmt.Println("Sending", dealID, "to user", userID)
		MarkDealAsSent(c, dealID, userID)
	}
}

func ShowUsersThatReceivedAllDeals(c *redis.Client, dealIDs []string) {
	users := c.SInter(dealIDs...).Val()
	fmt.Printf("%v received all of the deals: %v\n", users, dealIDs)
}

func ShowUsersThatReceivedAtLeastOneOfTheDeals(c *redis.Client, dealIDs []string) {
	users := c.SUnion(dealIDs...).Val()
	fmt.Printf("%v received at least one of the deals: %v\n", users, dealIDs)
}

func main() {
	c, err := NewClient()
	if err != nil {
		panic(err)
	}

	MarkDealAsSent(c, "deal:1", "user:1")
	MarkDealAsSent(c, "deal:1", "user:2")
	MarkDealAsSent(c, "deal:2", "user:1")
	MarkDealAsSent(c, "deal:2", "user:3")

	SendDealIfNotSent(c, "deal:1", "user:1")
	SendDealIfNotSent(c, "deal:1", "user:2")
	SendDealIfNotSent(c, "deal:1", "user:3")

	ShowUsersThatReceivedAllDeals(c, []string{"deal:1", "deal:2"})
	ShowUsersThatReceivedAtLeastOneOfTheDeals(c, []string{"deal:1", "deal:2"})
}
