package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func func1(){
	fmt.Println("func1 called")
	wg.Done()
}

func func2(){
	fmt.Println("func2 called")
	wg.Done()
}
func main() {
	fmt.Println("main")
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Add(2)
	go func1()
	fmt.Println("after func1")
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	go func2()
	fmt.Println("after func2")
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Done")
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}
