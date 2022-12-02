package main

// package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"reflect"
	// "runtime"
	"strings"

	// "github.com/gin-gonic/gin"
	// "github.com/go-playground/locales/sv"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nats-io/nats.go"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
	// "github.com/onsi/ginkgo/config"
	// "golang.org/x/net/http2"
	// "golang.org/x/net/http2/h2c"
	//  util "github.com/franklynobleC/dictionaryAPIGrpc/util"
	se "github.com/franklynobleC/dictionaryAPIGrpc/proto"
	// s"github.com/franklynobleC/dictionaryAPIGrpc/proto/pb"google.golang.org/grpc"
)

// import "google.golang.org/genproto/googleapis/cloud/orchestration/airflow/service/v1"

// Dictioary Server Error

var (
	OpeningFileError = errors.New("error opening file")
	MarshallError    = errors.New("error encoding dictionary")
	WordNotFound     = errors.New("keyword not found, try another key word")
	EnterKeyWord     = errors.New("enter a keyword")
	EmptyString      = errors.New("")
)

const (
	StreamName = "AllWORDS"
	// StreamSubject = "WORDS*"
)

// var word string

var (
	port = flag.Int("port", 8082, "Server port")
)

type Dyc interface{}
type Last struct {
	Word   string
	Meaning string
}

type server struct {
	se.UnimplementedEnglishDictionaryServiceServer
	jt nats.JetStreamContext
	// net.Listener
	// service pb.Repository
}

func NewServer() *server {
	return &server{}
}

func (serv *server) EnglishDictionarySearchWord(ctx context.Context, word *se.EnglishDictionarySearchWordRequest) (*se.EnglishDictionarySearchWordResponse, error) {

	// dd := &serv.Publish(StreamName, []byte(word.Word))
	// log.Println("receivedwords: %v")
	words := &se.EnglishDictionarySearchWordRequest{
		Word: string(word.GetWord()),
	}
	fmt.Print(words, "1st")

	strings.TrimSpace(strings.ToLower(words.GetWord()))
	fmt.Print(words, "2nd")

	if len(words.GetWord()) == 0 {

		return &se.EnglishDictionarySearchWordResponse{
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

	fmt.Print(reflect.ValueOf(Dyc).Len())

	//TODO: FOR NATS PUBLISHING

	jst, err := JetStreamInit()
	if err != nil {
		log.Fatal("cant connect to nats ", err.Error())
	}

	err = CreateStream(jst)
	if err != nil {
		log.Fatal("cant create stream ", err.Error())
	}

	// defer consumeWords(jst)
	var S = Last{
		
	Word:    word.Word,
		Meaning: Dyc[word.GetWord()],
	}
    //   for k,v := range S.meaning {

	//   }
	_, kePresent := Dyc[words.GetWord()]

	if kePresent {
		fmt.Println(kePresent, "key present")
		//Publish(StreamName, []byte(Dyc[word.GetWord()]))
		// ExampleJetStream(  Dyc[word.GetWord()])

		stm, err := json.Marshal(S)
		if err != nil {
					
		log.Print(MarshallError, "from service A", err)
		}
		
		jst.Publish(StreamName, stm)

		log.Println("byt printed ACTUAL DATA  ", string(stm))

		return &se.EnglishDictionarySearchWordResponse{

			Words:   word.Word,
			Meaning: Dyc[word.GetWord()],
		}, nil
		// ExampleJetStream(Dyc[word.GetWord()])

	} else {
		return &se.EnglishDictionarySearchWordResponse{
			Words: string(""),
		}, WordNotFound
	}

	//  fmt.Print( string(valuebyte[]))

}

func main() {

	// ExampleJetStream(NewServer())

	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	err := se.RegisterEnglishDictionaryServiceHandlerServer(ctx, grpcMux, NewServer())

	if err != nil {
		log.Fatal("can not register handler Server", err)
	}

	mux := http.NewServeMux()

	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", ":5000")

	if err != nil {
		log.Fatal("can not create listener", err)
	}

	log.Println("http Gateway Server is being Started", listener.Addr().String())

	err = http.Serve(listener, mux)

	if err != nil {
		log.Fatal("can not start grpc server", err)
	}
}

func JetStreamInit() (nats.JetStreamContext, error) {

	//connect to NATS

	nc, err := nats.Connect("nats://0.0.0.0:4222")

	if err != nil {
		return nil, errors.New("coudl not connect to Nats")
	}

	log.Println("connected to Jetstream", nc.ConnectedAddr())
	//create JetStream Context

	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))

	if err != nil {
		log.Println("could not publish to Jestream")
		return nil, err
	}
	log.Println("successfully published JetStream")
	return js, nil

}

func CreateStream(jetStream nats.JetStreamContext) error {

	stream, err := jetStream.StreamInfo(StreamName)
	// stream  not found ,create it

	if stream == nil {
		log.Printf("creating stream: %s\n", StreamName)

		_, err = jetStream.AddStream(
			&nats.StreamConfig{
				Name: StreamName,
				// Subjects: []string{"test"},
			},
		)

		// fmt.Print(ack)
		if err != nil {
			log.Println("could not add  stream")
			return err
		}

	}
	return nil

}

func consumeWords(js nats.JetStreamContext) {
	_, err := js.Subscribe(StreamName, func(m *nats.Msg) {
		err := m.Ack()

		if err != nil {
			log.Println("Unable to Ack", err)
			return
		}

		fmt.Println(string(m.Data))

		log.Println("Successfully consumed")

		//		log.Printf("Consumer  =>  Subject: %s  -  ID:%s  -  Author: %s  -  Rating:%d\n", m.Subject, review.Id, review.Author, review.Rating)

		// send answer via JetStream using another subject if you need
		// js.Publish(config.SubjectNameReviewAnswered, []byte(review.Id))
	})

	if err != nil {
		log.Println("Subscribe failed")
		return
	}
}
