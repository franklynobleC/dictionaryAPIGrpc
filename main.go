package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	// "unicod
	 ed  "utill"
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
	WordsNotFound = errors.New("keyword not found, try another key word")
)

var Meaning = make(map[string]string)

type Dictionary struct {
	words map[string]search
}

type search interface{}

var Dyc interface{}

func main() {

	result := ed.SearchWord("god")

	fmt.Print(result)

	// 	if err != nil {
	// 		fmt.Println("in marshalling error")
	// 	}

	// 	jsonfile, err := os.Open("dictionary.json")

	// 	if err != nil {
	// 		fmt.Println(err.Error())

	// 	}

	// 	fmt.Print("successfully opened")

	// 	valuebyte, _ := ioutil.ReadAll(jsonfile)

	// 	err = json.Unmarshal(valuebyte, &Dyc)

	// 	for i, _ := range string(valuebyte) {
	// 		strconv.Itoa(int(valuebyte[i]))
	// 		i++
	// 		// fmt.Print(type(i))
	// 		// fmt.Println(types.)
	// 	}
	// 	// fmt.Print(Dyc)

	// 	fmt.Print(string(valuebyte[:3]))

	// 	fmt.Print(reflect.ValueOf(Dyc).Len())
	//   map1 := Dyc.(map[string] interface{})

	// 	fmt.Println(map1["okay"])

	// 	// if Dyc.Kind() == reflect.Map{

	// 	// }

	// 	//  fmt.Print( string(valuebyte[]))
	// 	// EnterWord()

	// 	Meaning["hello"] = "greetings human way"
	// 	Meaning["hammer"] = "tool for nailing into a device"
	// 	Meaning["James"] = "human name"
	// 	Meaning["computer"] = "electronice device that processes data"
}

// Take a slice of keys, say band names that are similar
// http://www.tonedeaf.com.au/412720/38-bands-annoyingly-similar-names.htm
// wordsToTest := []string{"King Gizzard", "The Lizard Wizard", "Lizzard Wizzard"}

// // Choose a set of bag sizes, more is more accurate but slower
// bagSizes := []int{2}

// // Create a closestmatch object
// // cm := closestmatch.New(wordsToTest, bagSizes)

// fmt.Println(cm.Closest("kind gizard"))
// // returns 'King Gizzard'

// fmt.Println(cm.ClosestN("kind gizard",3))
// returns [King Gizzard Lizzard Wizzard The Lizard Wizard]

// fmt.Println(searchWords("hello"))

//  result :=  searchWords(ds.mean["hello"])

// 	for k, v := range ds.mean {
// 		strings.ToUpper(k)
// 		if k == "hello" {
// 			fmt.Print(k, " "+v)
// 		}
// 		break
// 	}
// }

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
