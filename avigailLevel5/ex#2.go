package main

import "fmt"

type person2 struct {
	firstName       string
	lastName        string
	iceCreamFlavors []string
}

func main() {
	p1 := person2{
		firstName:       "6525",
		lastName:        "last name 1",
		iceCreamFlavors: []string{"asas", "asas"},
	}
	p2 := person2{
		firstName:       "5656",
		lastName:        "bbb",
		iceCreamFlavors: []string{"asas", "asas"},
	}
	m := map[string]person2{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		fmt.Print("iceCreamFlavors: ")
		for _, fla := range v.iceCreamFlavors {
			fmt.Print(fla, ", ")
		}
		fmt.Println()
	}

}
