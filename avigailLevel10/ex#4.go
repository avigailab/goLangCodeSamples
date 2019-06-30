package main

import (
	"fmt"
)

func main() {
	q1 := make(chan int)
	c1 := gen1(q1)

	receive1(c1, q1)

	fmt.Println("about to exit")
}

func receive1(c <-chan int, q <-chan int)  {
	//infinite loop
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return

		}
	}

}
func gen1(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}
