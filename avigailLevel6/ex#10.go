package main

import "fmt"

func func4() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	f1 := func4()
	fmt.Println("F1")
	fmt.Println(f1())
	fmt.Println(f1())
	f2 := func4()
	fmt.Println("F2")
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println("F1")
	fmt.Println(f1())
	fmt.Println(f1())

}
