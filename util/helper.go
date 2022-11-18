package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

var (
	MarshallError = errors.New("error decoding dictionary")
	WordNotFound  = errors.New("keyword not found, try another key word")
	EnterKeyWord  = errors.New("enter a keyword")
)

//TODO: remove debug errors in this package

type Dyc interface{}

//TODO: Add errors if words not found

func SearchWord(word string) (string, error) {

	word = strings.TrimSpace(strings.ToLower(word))

	if len(word) == 0 {
		return "", EnterKeyWord
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
		return "", MarshallError
	}

	fmt.Print(reflect.ValueOf(Dyc).Len())

	_, kePresent := Dyc[word]
	if kePresent {
		fmt.Println(kePresent)
		return Dyc[word], nil
	} else {
		return "", WordNotFound
	}

	//  fmt.Print( string(valuebyte[]))

}
