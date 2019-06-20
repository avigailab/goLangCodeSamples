package main

import "fmt"

func main() {
	myArray := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//fmt.Println(myArray)
	for n := range myArray {
		fmt.Printf("number: %v, type; %T\n", n, n)
	}
	fmt.Printf("array type: %T", myArray)

}
