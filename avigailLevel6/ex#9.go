package main

import "fmt"

func func3() func() int {
	return func() int {
		fmt.Println("Anonymous func")
		return 3
	}
}
func callback(f func() int, i int) int{
	a := f()
	return a*i
}
func main() {
	s := callback(func3() ,3)
	fmt.Println(s)
}
