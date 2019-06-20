package main

import "fmt"

type ex5Type int
var x3 ex5Type
var y3 int
func main() {
	//a. print out the value of the variable “x”
	fmt.Println(x3)
	// b. print out the type of the variable “x”
	fmt.Printf("%T\n", x3)
	// c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
	x3 = 42
	// d. print out the value of the variable “x
	fmt.Println(x3)
	y3 = int(x3)
	fmt.Println(y3)
	fmt.Printf("%T\n", y3)
}