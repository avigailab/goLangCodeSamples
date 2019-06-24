package main

import "fmt"

func foo(a int) int{
	return a
}

func bar(a int, b string) (int, string){
	return a, b
}

func main() {
	a, b := bar(20, "asas")
	fmt.Println(foo(30),a, b)
}
