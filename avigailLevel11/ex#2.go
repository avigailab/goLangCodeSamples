package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person1 struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person1{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("Error %v in marshalling %v", err, a)
	}
	return bs,err
}
