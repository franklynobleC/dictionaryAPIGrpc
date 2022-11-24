package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6658",
	})
	if err != nil {
		log.Fatal("connection NOT successful")
	}

	defer client.Close()

	producer, err := client.CreateProducer(
		pulsar.ProducerOptions{
			Topic: "topic-1",
		})

	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	ctx := context.Background()

	for i := 0; i < 10; i++ {
		if msgId, err := producer.Send(ctx, &pulsar.ProducerMessage{
			Payload: []byte(fmt.Sprint("hello", i)),
		}); err != nil {
			log.Fatal(err)
		} else {
			log.Println("Published message:", msgId)
		}
	}

}
