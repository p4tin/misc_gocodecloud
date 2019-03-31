package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()
	proj := "ecomm-dev-poc"
	client, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}
	publish(client, "cache_invalidation_topic", "This is a test")
}

func publish(client *pubsub.Client, topic, msg string) error {
	ctx := context.Background()
	// [START pubsub_publish]
	// [START pubsub_quickstart_publisher]
	t := client.Topic(topic)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	// [END pubsub_publish]
	// [END pubsub_quickstart_publisher]
	return nil
}
