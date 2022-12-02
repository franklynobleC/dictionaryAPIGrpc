package main

import (
	// "context"
	// "context"s
	// "errors"
	"context"
	"encoding/json"
	// "encoding/json"
	"fmt"
	"log"
	"net"
	"net/http" //would uncomment
	// "os/user"
	"time"
	// "sync"
	// "context"

	hs "github.com/franklynobleC/dictionaryAPIGrpc/grpcServiceB/history/proto"
	// "google.golang.org/grpc"
	// "github.com/go-playground/locales/id"
	// "github.com/go-playground/locales/id_ID"
	// "github.com/grpc-ecosystem/grpc-gateway/runtime"
	// "golang.org/x/text/unicode/rangetable"
	// "github.com/go-playground/locales/asa"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime" //would uncomment
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
	// "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	"github.com/nats-io/nats.go"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

type LastFromSubscribe struct {
	Word    string
	Meaning string
}

func (sv *server) DictionaryHistory(context.Context, *hs.DictionaryHistoryRequest) (*hs.DictionaryHistoryResponse, error) {
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
		ID      primitive.ObjectID `bson:"_id"`
		Word    string             `bson:"word"`
		Meaning string             `bson:"meaning"`
	}{}

	// TD := []struct {
	// 	id   string
	// 	word string
	// }{}
	// Gresp := hs.DictionaryHistoryResponse{}
	// defer cur.Close(context.Background())
	var Gresp *hs.DictionaryHistoryResponse
	for cur.Next(context.Background()) {

		err = cur.Decode(&res)

		if err != nil {
			log.Println("can not get result")
		}
		fmt.Print("\n")
		fmt.Print("------------------------------------------------")
		fmt.Println(res.ID)
		// tests := res.m
		// keys := []string{}

		fmt.Println("FROM MARSHALLING")
		Gresp = &hs.DictionaryHistoryResponse{
			History: []*hs.History{
				{

					Id:      res.ID.Hex(),
					Word:    res.Word,
					Meaning: res.Meaning,
				}, 
			},
			// Sleep for a little bit.
		}

	}
	return Gresp, nil
}

// fmt.Println(res.ID)
// tests := res.m
// keys := []string{}

// fmt.Println("FROM MARSHALLING 22")
// return Gresp, nil
// }
// 	return &hs.DictionaryHistoryResponse{
// 		History: []*hs.History{
// 			  &*SDS[

// 			  ]
// 		},
// 	},

// return nil, nil

// }

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

	// nc.InMsgs

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
	log.Print("from metadata", msg.Subject)

	var Sub LastFromSubscribe

	err = json.Unmarshal(msg.Data, &Sub)

	if err != nil {
		log.Println("ERROR UNMARSHALLING FROM SERVICE B", err)
	}

	// for _, v := range Sub {

	// log.Println(v.Word, sub.Meaning)
	// Sub = append(Sub, v)

	// wordTO := string(msg.Data)
	// va(msg.Ddata
	// res, err := json.Marshal(msg.Data)

	// WordM := make(map[string]interface{})
	// var ccc string

	// word1 := struct{
	// 	  word string
	// 	  meaning string
	// }{
	// 	 word:string(msg.Data),
	// 	meaning: string(msg.Data),
	// }

	// sds := string(msg.Data)
	// var v byte
	// for _, v = range msg.Data {
	// 	fmt.Print(v)

	// }

	//  word1 = struct{word string; meaning string}{
	// 	 word:string(msg.Data),
	// 	meaning: string(msg.Data),
	//  }
	// var wordMean string
	// err = json.Unmarshal(msg.Data, &ccc)
	// ss := string(msgData)
	log.Printf("Word: Meaning %s", string(msg.Data))
	// smg := string(msg.Data)

	// msg.Metadata().
	nn := bson.D{{Key: "word", Value: Sub.Word}, {Key: "meaning", Value: Sub.Meaning}}

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

}

//fmt.Println(consumeWords(jst))

func main() {

	//   opts := []grpc.ServerOption{}

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
