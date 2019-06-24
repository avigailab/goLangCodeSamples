package main

import "fmt"

func foo1(n ...int) int{
	sum := 0
	for _,v := range n{
		sum +=v
	}
	return sum
}

func bar1(n []int) int{
	sum := 0
	for _,v := range n{
		sum +=v
	}
	return sum
}
func main() {
	fmt.Println(foo1(30,30,5))
	a := []int{30,30,30}
	fmt.Println(bar1(a))
	fmt.Println(foo1(a...))
}
