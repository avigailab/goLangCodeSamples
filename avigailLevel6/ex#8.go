package main

import "fmt"

func func2() func() {
	return func() {
		fmt.Println("Anonymous func")
	}
}
func main() {
	s := func2()
	s();
}
