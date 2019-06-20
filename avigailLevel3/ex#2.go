package main

import "fmt"

func main() {
	for i := 61; i <= 122; i++ {
		for j :=0; j < 3; j++ {
			fmt.Printf("%#U\t\n",i)
		}
	}
}
