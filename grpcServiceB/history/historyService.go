package main

import (
	// "context"
	// "context"s
	// "errors"
	"context"
	// "encoding/json"
	"net"
	"net/http"
	// "strings"
	// "encoding/json"
	"fmt"
	"log"
	// "os/user"
	"time"
	// "sync"
	// "context"

	hs "github.com/franklynobleC/dictionaryAPIGrpc/grpcServiceB/history/proto"
	// "github.com/go-playground/locales/asa"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
	// "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	"github.com/nats-io/nats.go"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type server struct {
	hs.UnimplementedDictionaryHistoryServiceServer
	// database *mongo.Client
}

var (
	StreamName = "AllWORDS"
)

func NewServer() *server {
	return &server{}
}

func (sv *server) DictionaryHistory(ctx context.Context, req *hs.DictionaryHistoryRequest) (*hs.DictionaryHistoryResponse, error) {
	// log.Println("Get all Data from Database", hs.	)// err := sv.

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://golangdb:testdb@cluster0.qz143pi.mongodb.net/?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal("could not connect to mongo Db")
	}
	fmt.Print("database connected successfully")

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err.Error())
	}

	wordDictionary := client.Database("borderlesshq").Collection("dictionary")

	fmt.Print("database created", wordDictionary.Database())
	fmt.Print("database created", wordDictionary.Database())

	cur, err := wordDictionary.Find(context.Background(), bson.D{})

	if err != nil {
		log.Println("returned error from getting data", err.Error())
	}

	res := struct {
		m map[string]string
	}{}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {

		err := cur.Decode(&res.m)

		if err != nil {
			log.Println("can not get result")
		}


	fmt.Print("\n")
	fmt.Print("------------------------------------------------")
	fmt.Println(res.m)
	// tests := res.m
	keys := []string{}

	for k, v := range res.m {
		keys = append(keys, string(k +""+v))
	}

	fmt.Println("FROM MARSHALLING")
	return &hs.DictionaryHistoryResponse{
		History: keys,
		// Sleep for a little bit..

	}, nil
	}

	return nil, nil
}

//  }

// func (sv *service) GetAllWords(ctx context.Context, mongo *mongo.Client, data string) (string, error) {

// res , err := mongo.
// }

// 	//TODO: FOR NATS PUBLISHING

// func(word *mongo.Client) (*mongo.)

func subScribeAndWrite() {

	// 	//TODO: FOR SUBSCRIBING PUBLISHING

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://golangdb:testdb@cluster0.qz143pi.mongodb.net/?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal("could not connect to mongo Db")
	}
	fmt.Print("database connected successfully")

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err.Error())
	}

	wordDictionary := client.Database("borderlesshq").Collection("dictionary")

	fmt.Print("database created", wordDictionary.Database())

	//TODO: NATS FUNCTIOINS TO SUBSCRUBE
	nc, err := nats.Connect("nats://0.0.0.0:4222")

	if err != nil {
		log.Println("coudl not connect to Nats", err.Error())
	}

	log.Println("connected to Jetstream", nc.ConnectedAddr())

	//subscribe
	sub, err := nc.SubscribeSync(StreamName)

	if err != nil {
		log.Print("error subscribing", err)
	}
	//wait for a  message
	//wait for this number of seconds to get the using  this time out
	msg, err := sub.NextMsg(50 * time.Second)

	if err != nil {
		log.Fatal(err.Error())
	}

	//use  the response
	log.Printf("Reply: %s", msg.Data)

	// wordTO := string(msg.Data)
	// va(msg.Ddata
	// res, err := json.Marshal(msg.Data)

	// WordM := make(map[string]interface{})
	// var ccc string
	sds := string(msg.Data)
	// var wordMean string
	// err = json.Unmarshal(msg.Data, &ccc)
	nn := bson.D{{Key: "words", Value: sds}}

	if err != nil {
		log.Print("can not unmarshal")
	}

	words, err := wordDictionary.InsertOne(context.TODO(), nn)

	if err != nil {
		log.Print("could not insert data", err.Error())
	}
	//else diplay the id of the newly inserted ID
	fmt.Println(words.InsertedID)

	fil, err := wordDictionary.Find(context.TODO(), nn)

	// Ok := fil.Next(context.TODO())

	defer fil.Close(context.Background())
	// fmt.Println(fil.Next(context.TODO()))
	//  fmt.Print(fil.Decode(fil))

	for fil.Next(context.Background()) {

		result := struct {
			m map[string]string
		}{}

		err := fil.Decode(&result)

		if err != nil {
			log.Fatal(err.Error(), "decoding data")
		}

	}

	//fmt.Println(consumeWords(jst))

}

func main() {
	defer subScribeAndWrite()

	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	err := hs.RegisterDictionaryHistoryServiceHandlerServer(ctx, grpcMux, NewServer())

	if err != nil {
		log.Fatal("can not register handler Server", err)
	}

	mux := http.NewServeMux()

	//
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

// func consumeWords(js nats.JetStreamContext) {

// 	_, err := js.Subscribe(StreamName, func(m *nats.Msg) {
// 		err := m.Ack()

// 		if err != nil {
// 			log.Println("Unable to Ack", err)
// 			return
// 		}

// 		js.Consumers("consumer")
// 		fmt.Println(string(m.Data))

// 		log.Println("Successfully consumed From Service BB")

// 		//
// 		// log.Printf("Consumer  =>  Subject: %s  -  ID:%s  -  Author: %s  -  Rating:%d\n", m.Subject, review.Id, review.Author, review.Rating)
// 		// send answer via JetStream using another subject if you need
// 		// js.Publish(config.SubjectNameReviewAnswered, []byte(review.Id))
// 	})

// 	if err != nil {
// 		log.Println("Subscribe failed")
// 		return
// 	}

// }
