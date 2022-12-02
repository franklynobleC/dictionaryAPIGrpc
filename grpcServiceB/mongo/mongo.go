package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/franklynobleC/dictionaryAPIGrpc/grpcServiceB/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// // Handler ...
// type Handler struct {
// 	mgo.Session
// }

// GetTodos ...
func GetHistory() ([]models.Histiory, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://golangdb:testdb@cluster0.qz143pi.mongodb.net/?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal("could not connect to mongo Db")
	}
	fmt.Print("database connected successfully")

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err.Error())
	}
	// defer cur.Close(context.Background())
	wordDictionary := client.Database("borderlesshq").Collection("dictionary")

	fmt.Print("database created", wordDictionary.Database())

	cur, err := wordDictionary.Find(context.Background(), bson.D{})

	if err != nil {
		log.Println("returned error from getting data", err.Error())
	}

	if err != nil {
		log.Println("returned error from getting data", err.Error())
	}

	histories := []models.Histiory{}
	for cur.Next(context.Background()) {

		err := cur.Decode(&histories)

		if err != nil {
			log.Println("can not get result")
		}

		// "borderlesshq").Collection("dictionary")
		return histories, err
	}
	return nil, err

}
