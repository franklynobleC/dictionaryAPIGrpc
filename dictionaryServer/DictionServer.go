package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"reflect"
	"strings"

	//  util "github.com/franklynobleC/dictionaryAPIGrpc/util"
	pb "github.com/franklynobleC/dictionaryAPIGrpc/pb"
	"google.golang.org/grpc"
)

// import "google.golang.org/genproto/googleapis/cloud/orchestration/airflow/service/v1"

// Dictioary Server Error

var (
	OpeningFileError = errors.New("error opening file")
	MarshallError    = errors.New("error decoding dictionary")
	WordNotFound     = errors.New("keyword not found, try another key word")
	EnterKeyWord     = errors.New("enter a keyword")
	EmptyString      = errors.New("")
)

var (
	port = flag.Int("port", 5000, "Server port")
)

type Dyc interface{}

type server struct {
	pb.UnimplementedEnglishDictionaryServer
	// service pb.Repository
}

// func (serv *server)returnWords (*service.ReturnWords, error) {

// }

func (serv *server) SearchWords(ctx context.Context, word *pb.Wordrequest) (*pb.WordResponse, error) {

	// log.Println("receivedwords: %v")
	words := &pb.Wordrequest{
		Word: string(word.GetWord()),
	}
	fmt.Print(words, "1st")

	strings.TrimSpace(strings.ToLower(words.GetWord()))
	fmt.Print(words, "2nd")

	if len(words.GetWord()) == 0 {

		return &pb.WordResponse{
			Words: EmptyString.Error(),
		}, EnterKeyWord
	}
	fmt.Print("before opening", words.GetWord())

	var Dyc map[string]string

	jsonfile, err := os.Open("dictionary.json")

	if err != nil {
		fmt.Println(OpeningFileError)

	}

	fmt.Print("successfully opened")

	valuebyte, _ := ioutil.ReadAll(jsonfile)

	err = json.Unmarshal(valuebyte, &Dyc)

	if err != nil {
		return &pb.WordResponse{
			Words: fmt.Sprint(Dyc[EmptyString.Error()]),
		}, MarshallError
	}

	fmt.Print(reflect.ValueOf(Dyc).Len())

	_, kePresent := Dyc[words.GetWord()]
	if kePresent {
		fmt.Println(kePresent, "key present")

		// wd := &pb.WordResponse{
		// 	Words: fmt.Sprint(Dyc[word1]),
		// }
		return &pb.WordResponse{
			Words: Dyc[word.GetWord()],
		}, nil
	} else {
		return &pb.WordResponse{
			Words: string(""),
		}, WordNotFound
	}

	//  fmt.Print( string(valuebyte[]))

}

func main() {
	flag.Parse()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("errors %v failed to liten in port %v", err, *port)
	}

	serv := grpc.NewServer()

	pb.RegisterEnglishDictionaryServer(serv, &server{})
	log.Printf("server listening at %v", listen.Addr())

	if err := serv.Serve(listen); err != nil {
		log.Fatalf("failed to Serve %v", err)
	}

}
