package main

import (
	"fmt"
)

func main() {
	//1
	c1 := gen2()
	receive2(c1)
 	//2
	c3 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c3 <- i
		}
		close(c3)
	}()

	for v := range c3 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func receive2(c <-chan int)  {
	for {
		v , ok:= <-c
		if ok {
			fmt.Println(v, ok)
		} else {
			return
		}
	}

}
func gen2() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
