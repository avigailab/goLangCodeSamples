package main

import "fmt"

func main() {
	var x int = 13
	fmt.Printf("%d\t,%b\t,%x\t\n",x,x,x)
	y := x << 1
	fmt.Printf("%d\t,%b\t,%x\t",y,y,y)

}
