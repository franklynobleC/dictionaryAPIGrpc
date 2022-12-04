package main

import (
	"context"
	"encoding/json"
	"os"

	// "encoding/json"
	"fmt"
	"log"
	"net"
	"net/http" //would uncomment

	// "os/user"
	"time"

	hs "github.com/franklynobleC/dictionaryAPIGrpc/grpcServiceB/history/proto" //Service B protoFile  is  in This Directory
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"                        //would uncomment
	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

type server struct {
	hs.UnimplementedDictionaryHistoryServiceServer
	S *mongo.Client
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

	dictionaryClient, err := ConnectMongo()

	fmt.Print("database created", dictionaryClient.Database())
	fmt.Print("database created", dictionaryClient.Database())

	cur, err := dictionaryClient.Find(context.Background(), bson.D{})

	if err != nil {
		log.Println("returned error from getting data", err.Error())
	}

	list := []*hs.History{}
	for cur.Next(context.Background()) {

		lt := new(hs.History)

		err = cur.Decode(&lt)

		if err != nil {
			log.Println("can not get result")
		}
		fmt.Print("\n")
		fmt.Print("------------------------------------------------")
		fmt.Println(lt.Word)
		// tests := res.m
		// keys := []string{}
		list = append(list, lt)
	}

	fmt.Println("FROM MARSHALLING")

	return &hs.DictionaryHistoryResponse{
		Histories: list,
	}, nil

}

//subscribe NATS to a topic and write to Database
func subScribeAndWrite() {

	// 	//TODO: connect to Database and get Database Client

	wordDictionary, err := ConnectMongo()

	if err != nil {
		log.Fatal("could not connect to mongo Db")
	}
	fmt.Print("database connected successfully")

	fmt.Print("database created", wordDictionary.Database())

	//TODO: NATS CONNECTION
	//subscribe to natsTopic
	nc, err := ConnectToNats()
	if err != nil {
		log.Println("could not connet to nats", err)
	}

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

func main() {

	// go subScribeAndWrite()

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

	listener, err := net.Listen("tcp", ":6000")

	if err != nil {
		log.Fatal("can not create listener", err)
	}

	log.Println("http Gateway Server is being Started", listener.Addr().String())
	subScribeAndWrite()

	err = http.Serve(listener, mux)

	if err != nil {
		log.Fatal("can not start grpc server", err)

	}

}

///
//Connet To Mongo Db and Return db Client
func ConnectMongo() (*mongo.Collection, error) {

	// Get DB data from .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("could not Load .env file")
	}

	opts := options.Client().ApplyURI(os.Getenv("DB_URL"))

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatal("could not connect to mongo Db")
	}
	fmt.Print("database connected successfully FROM mongo func")

	wordDictionary := client.Database(os.Getenv("DB_NAME")).Collection("dictionary")

	fmt.Print("database created", wordDictionary.Database())

	return wordDictionary, nil
}

func ConnectToNats() (*nats.Conn, error) {

	nc, err := nats.Connect(os.Getenv("JESTREAM_URL"))

	if err != nil {
		log.Println("coudl not connect to Nats", err.Error())
	}

	log.Println("connected to Jetstream", nc.ConnectedAddr())

	//subscribe

	return nc, nil
}
