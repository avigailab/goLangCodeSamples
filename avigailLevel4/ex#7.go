package main

import "fmt"

func main() {
	a := []string{"James", "Bond", "Shaken, not stirred" }
	b := []string{"Miss", "Moneypenny", "Helloooooo, James." }

	myArray := [][]string{a, b}
	//fmt.Println(myArray)
	for i := range myArray {
		fmt.Println("record ", i)
		for _ , v := range myArray[i] {
			fmt.Println(v)
		}
	}

}
