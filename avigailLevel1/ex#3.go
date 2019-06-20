package main

import "fmt"

var x1 int = 42
var y1 string = "avigail"
var z1 bool = true

func main() {
	s := fmt.Sprintf("%v , %s, %t\n" ,x1, y1, z1)
	fmt.Println(s)
}