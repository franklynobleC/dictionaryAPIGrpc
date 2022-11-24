package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
	// "github.com/Comcast/pulsar-client-go"
	"github.com/apache/pulsar/pulsar-function-go/pf"
)

func PublishFunc(ctx context.Context, in []byte) error {
	fctx, ok := pf.FromContext(ctx)

	if !ok {
		return errors.New("get Go functions Context error")
	}

	publishTopic := "publish-topic"
	output := append(in, 110)

	producer := fctx.NewOutputMessage(publishTopic)
	msgID, err := producer.Send(ctx, &pulsar.ProducerMessage{
		Payload: output,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("The output message ID is: %+v", msgID)
	return nil
}

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{URL: "pulsar://localhost:6650"})

	if err != nil {
		log.Fatal(errors.New("connection NOT successful"))
	}
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "topic-1",
		SubscriptionName: "my-sub",
		Type:             pulsar.KeyShared,
	})

	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	for i := 0; i < 10; i++ {
		msg, err := consumer.Receive(context.Background())

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
			msg.ID(), string(msg.Payload()))

		consumer.Ack(msg)
	}

	if err := consumer.Unsubscribe(); err != nil {
		log.Fatal(err)
	}

}
