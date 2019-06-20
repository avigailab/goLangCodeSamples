package main

import "fmt"

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	//fmt.Println(myArray)
	for n := range myArray {
		fmt.Printf("number: %v, type; %T\n", n, n)
	}
	fmt.Printf("array type: %T", myArray)

}
