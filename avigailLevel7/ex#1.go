package main

import "fmt"

func main() {
	n := 30
	p := &n
	fmt.Println(n, p, *p)
	*p = 34
	fmt.Println(n, p, *p)
}
