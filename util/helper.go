package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
)

var (
	MarshallError = errors.New("error decoding dictionary")
)

type Dyc interface{}

//TODO: Add errors if words not found

func SearchWord(word string) string {
	var Dyc interface{}

	jsonfile, err := os.Open("dictionary.json")

	if err != nil {
		fmt.Println(err.Error())

	}

	fmt.Print("successfully opened")

	valuebyte, _ := ioutil.ReadAll(jsonfile)

	err = json.Unmarshal(valuebyte, &Dyc)

	if err != nil {
		return fmt.Sprint(MarshallError)
	}

	for i, _ := range string(valuebyte) {
		strconv.Itoa(int(valuebyte[i]))
		
		// fmt.Print(type(i))
		// fmt.Println(types.)
	}
	// fmt.Print(Dyc)

	// fmt.Print(string(valuebyte[:3]))

	fmt.Print(reflect.ValueOf(Dyc).Len())
	map1 := Dyc.(map[string]interface{})

	return fmt.Sprint(map1[word])

	// if Dyc.Kind() == reflect.Map{

	// }

	//  fmt.Print( string(valuebyte[]))

}
