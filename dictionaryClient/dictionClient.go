package main

import (
	"context"
	"flag"
	"log"
	"time"

	// "github.com/apache/pulsar-client-go/integration-tests/pb"
	pb "github.com/franklynobleC/dictionaryAPIGrpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultword = "world"
)

var (
	addr = flag.String("add", "localhost:5000", "address to connect to")
	word = flag.String("word", defaultword, "word to search for")
)

func main() {

	flag.Parse()

	//Set up a connection to the server
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewEnglishDictionaryClient(conn)
   //contact the server and  print out its response
   ctx, cancel := context.WithTimeout(context.Background(), time.Second)  

     defer cancel()
    r, err := c.SearchWords(ctx, &pb.Wordrequest{
		Word: *word,
	})  
	if  err != nil {
		log.Fatalf("could not search word: %v", err)
	}
	log.Printf("word meaning:%s", r.GetWords())

}