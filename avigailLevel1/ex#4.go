package main

import "fmt"

type ex4Type int
var x2 ex4Type
func main() {
	//a. print out the value of the variable “x”
	fmt.Println(x2)
	// b. print out the type of the variable “x”
	fmt.Printf("%T\n", x2)
	// c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
	x2 = 42
	// d. print out the value of the variable “x
	fmt.Println(x2)
}