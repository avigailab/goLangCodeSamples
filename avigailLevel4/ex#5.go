package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//[42, 43, 44, 48, 49, 50, 51]
	var y []int
	y = append(y ,x[:3]...)
	y = append(y ,x[6:]...)
	fmt.Println(y)
	var z []int
	z = append(x[:3], x[6:]...)
	fmt.Println(z)
}
