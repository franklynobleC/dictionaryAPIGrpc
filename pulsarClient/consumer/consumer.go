package main

import (
	"context"
	"errors"
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
  cleint, err := pulsar.NewClient(pulsar.ClientOptions{}
)

}
