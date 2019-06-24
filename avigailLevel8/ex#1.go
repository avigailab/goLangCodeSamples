package main

import (
	"encoding/json"
	"fmt"
)

type user1 struct {
	First string
	Age   int
}

func main() {
	u1 := user1{
		First: "James",
		Age:   32,
	}

	u2 := user1{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user1{
		First: "M",
		Age:   54,
	}

	users := []user1{u1, u2, u3}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
