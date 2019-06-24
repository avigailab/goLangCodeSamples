package main

import "fmt"

func main() {
	s := func() {
		fmt.Println("Anonymous func")
	}
	s()
}
