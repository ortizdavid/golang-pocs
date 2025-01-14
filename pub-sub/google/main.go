package main

import (
	"context"
	"flag"
	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()

	projectId := flag.String("projectId", "", "ProjectcId")
	topicName := flag.String("topic", "", "Topic name")
	msg := flag.String("msg", "", "Message")
	flag.Parse()

	c, err := pubsub.NewClient(ctx, *projectId)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	topic := c.Topic(*topicName)
	exists, err := topic.Exists(ctx)
	if err != nil {
		panic(err)
	}

	if !exists {
		topic, err = c.CreateTopic(ctx, *topicName)
		if err != nil {
			panic(err)
		}
	}

	result := topic.Publish(ctx, &pubsub.Message{
		ID:              "",
		Data:            []byte(*msg),
	})
	_, err = result.Get(ctx)
	if err != nil {
		panic(err)
	}

}