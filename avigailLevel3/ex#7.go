package main

import "fmt"

func main() {
	i := 100
	if i == 10 {
		fmt.Println("i is 10")
	} else if i > 10 {
		fmt.Println("i bigger then 10", i)
	} else {
		fmt.Println("i is ", i)
	}
}
