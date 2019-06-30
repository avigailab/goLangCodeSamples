package main

import (
	"fmt"
)

func main() {
	// using func literal
	c := make(chan int)
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
	// using func literal
	c1 := make(chan int, 1)
	c1 <- 42
	fmt.Println(<-c1)
}
