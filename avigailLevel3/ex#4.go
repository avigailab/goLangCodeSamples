package main

import "fmt"

func main() {
	i := 1993
	for {
		fmt.Println(i)
		if i >= 2019 {
			break
		}
		i ++
	}
}
