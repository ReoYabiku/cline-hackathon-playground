package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pubsub := client.Subscribe(ctx, "ch00")
	_, err := pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	}

	log.Println("starting...")
	go func() {
		for {
			ch := pubsub.Channel()
			for msg := range ch {
				fmt.Println(msg.Channel, msg.Payload)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("stopping...")
	pubsub.Close()
}
