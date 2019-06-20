package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	iceCreamFlavors []string
}

func main() {
	p1 := person{
		firstName: "6525",
		lastName:  "last name 1",
		iceCreamFlavors: []string{"asas", "asas"},
	}
	p2 := person{
		firstName: "5656",
		lastName:  "bbb",
		iceCreamFlavors: []string{"asas", "asas"},
	}
	fmt.Println("p1")
	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for _, v := range p1.iceCreamFlavors {
		fmt.Println(v)
	}
	fmt.Println("p2")
	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)
	for _, v := range p2.iceCreamFlavors {
		fmt.Println(v)
	}

	all := []person{p1, p2}
	for i, p := range all {
		fmt.Println("p", i)
		fmt.Println(p.firstName)
		fmt.Println(p.lastName)
		fmt.Print("iceCreamFlavors: ")
		for _, v := range p.iceCreamFlavors {
			fmt.Print(v, ", ")
		}
		fmt.Println()
	}
}
