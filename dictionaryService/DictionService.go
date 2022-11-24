package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	pb "github.com/franklynobleC/dictionaryAPIGrpc/pb"
)

// import "google.golang.org/genproto/googleapis/cloud/orchestration/airflow/service/v1"

var (
	MarshallError = errors.New("error decoding dictionary")
	WordNotFound  = errors.New("keyword not found, try another key word")
	EnterKeyWord  = errors.New("enter a keyword")
)

type Dyc interface{}

type server struct {
	pb.UnimplementedEnglishDictionaryServer
	service pb.Repository
}

// func (serv *server)returnWords (*service.ReturnWords, error) {

// }

func (serv *server) SearchWords(ctx context.Context, word *pb.Wordrequest) (*pb.WordResponse, error) {

	// log.Println("receivedwords: %v")
	words := &pb.Wordrequest{
		Word: string(word.GetWord()),
	}

	word1 := strings.TrimSpace(strings.ToLower(string(fmt.Sprint(words))))

	if len(word1) == 0 {

		return &pb.WordResponse{
			Words: string(""),
		}, EnterKeyWord
	}

	var Dyc map[string]string

	jsonfile, err := os.Open("dictionary.json")

	if err != nil {
		fmt.Println(err.Error())

	}

	fmt.Print("successfully opened")

	valuebyte, _ := ioutil.ReadAll(jsonfile)

	err = json.Unmarshal(valuebyte, &Dyc)

	if err != nil {
		return &pb.WordResponse{
			Words: fmt.Sprint(Dyc[""]),
		}, MarshallError
	}

	fmt.Print(reflect.ValueOf(Dyc).Len())

	_, kePresent := Dyc[word1]
	if kePresent {
		fmt.Println(kePresent)

		wd := &pb.WordResponse{
			Words: fmt.Sprint(Dyc[word1]),
		}
		return wd, nil
	} else {
		return &pb.WordResponse{
			Words: fmt.Sprint(Dyc[""]),
		}, WordNotFound
	}

	//  fmt.Print( string(valuebyte[]))

}

func main() {

}
