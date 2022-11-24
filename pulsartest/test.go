package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:                        "pulsar://localhost:6650",
		OperationTimeout:           30 * time.Second,
		ConnectionTimeout:          30 * time.Second,
	
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	// fmt.Print(client.TopicPartitions("Tamp-1"))

	fmt.Print("pulsar client instanttiated successfully !!")

	producer, err := client.CreateProducer(
		pulsar.ProducerOptions{

			Topic: "topic-1",
		})

	if err != nil {
		log.Fatal(err)
	}

	msgID, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte("hello"),
	})

	fmt.Print(msgID)
	defer producer.Close()

	if err != nil {
		fmt.Println("Failed to publish message", err)
	}
	fmt.Println("Published message")

}
