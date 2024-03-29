package main

import (
	"fmt"
)

func receive (c <-chan int){
	for v := range c{
		fmt.Println(v)
	}
}
func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		c <- i
	}
	return c
}
