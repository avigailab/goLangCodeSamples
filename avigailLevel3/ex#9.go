package main

import "fmt"

func main() {
	var favSport string = "sport1"
	switch favSport{
	case "sport1":
		fmt.Println("sport1")
	case "sport2":
		fmt.Println("sport2")
	default:
		fmt.Println("default")
	}

}
