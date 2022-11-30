package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	// ed "util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type dyc interface{}

type ClosestMatch struct {
	SubstringSizes []int
	SubstringToID  map[string]map[uint32]struct{}
	ID             map[uint32]IDInfo
	mux            sync.Mutex
}

// IDInfo carries the information about the keys
type IDInfo struct {
	Key           string
	NumSubstrings int
}

var ()

var Meaning = make(map[string]string)

type Dictionary struct {
	words map[string]search
}

type search interface{}

// var Dyc interface{}

func main() {

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
	}

	// eer, _ := ed.SearchWord("MAN")

	// fmt.Print(eer)
	// fmt.Print(searchWord("Ok")

}

func findWords(word string) (string, error) {
	if len(word) == 0 {
		return "", errors.New("error no keyword found")

	}

	if tl := strings.HasPrefix(Meaning[word], Meaning["h"]); tl {

		// Meaning[word] = append(word,)
		if val, ok := Meaning[word]; ok {

			// fmt.Println(val)
			if ok {
				for k, v := range val {

					fmt.Println(k, " ", v)
					return "", nil
				}
			}

			return Meaning[word], nil

			// return val, nil

		}

	}
	return "", errors.New("error no key word")
}

//
