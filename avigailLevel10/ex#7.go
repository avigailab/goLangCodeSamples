package main

import (
	"fmt"
)

func main() {
	//1
	c1 := gen3()
	receive3(c1)
	fmt.Println("about to exit")
	//2

}

func receive3(c <-chan int)  {
	for {
		v , ok:= <-c
		if ok {
			fmt.Println(v, ok)
		} else {
			return
		}
	}

}
func gen3() <-chan int {
	c := make(chan int)
	for i := 0;i<10 ;i ++ {
		go func() {
			for j := i; j < 10*i; j++ {
				c <- j
			}
			close(c)
		}()
	}
	return c
}
