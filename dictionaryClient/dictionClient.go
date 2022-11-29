package main

import (
	"context"
	"flag"
	"log"
	"time"

	// "github.com/apache/pulsar-client-go/integration-tests/pb"
	se "github.com/franklynobleC/dictionaryAPIGrpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultword = "man"
)

var (
	addr = flag.String("add", "localhost:5000", "address to connect to")
	Word = flag.String("word", "man", "word to search for")
)

func main() {

	flag.Parse()

	//Set up a connection to the server
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := se.NewEnglishDictionaryServiceClient(conn)
	//contact the server and  print out its response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()
	r, err := c.EnglishDictionarySearchWord(ctx, &se.EnglishDictionarySearchWordRequest{
		Word: *Word,
	})
	if err != nil {
		log.Fatalf("could not search word: %v", err.Error())
	}
	log.Printf("word meaning:%s", r.GetWords())

}
