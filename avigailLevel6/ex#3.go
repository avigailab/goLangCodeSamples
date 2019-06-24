package main

import "fmt"

func func1() {
	defer func() {
		fmt.Println("func1 deffer ran")
	}()
	fmt.Println("func1 run")

}

func main() {
	defer func1()
	fmt.Println("main end")
}
