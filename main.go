package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	// ed "util"

	 ed "github.com/franklynobleC/dictionaryAPIGrpc/util"
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

var (
	
)

var Meaning = make(map[string]string)

type Dictionary struct {
	words map[string]search
}

type search interface{}

// var Dyc interface{}

func main() {
	eer, _ := ed.SearchWord("MAN")

	fmt.Print(eer)
	// fmt.Print(searchWord("Ok"))
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

			return val, nil

		}

	}
	return "", errors.New("error no key word")
}

//
