package main

import (
	"fmt"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{URL: "pulsar://localhost:6650"})
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	channel := make(chan pulsar.ConsumerMessage, 100)

	options := pulsar.ConsumerOptions{
		Topic:            "topic-1",
		SubscriptionName: "my-subscription",
		Type:             pulsar.Shared,
	}

	options.MessageChannel = channel

	consumer, err := client.Subscribe(options)
	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close()

	// Receive messages from channel. The channel returns a struct which contains message and the consumer from where
	// the message was received. It's not necessary here since we have 1 single consumer, but the channel could be
	// shared across multiple consumers as well
	for cm := range channel {
		msg := cm.Message
		fmt.Printf("Received message  msgId: %v -- content: '%s'\n",
			msg.ID(), string(msg.Payload()))

		consumer.Ack(msg)
	}
}
