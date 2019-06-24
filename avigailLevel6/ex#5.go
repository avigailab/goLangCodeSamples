package main

import (
	"fmt"
	"math"
)

type square struct{
	l int
	w int
}

type circle struct{
	r int
}

type shape interface {
	area()
}

func (c circle) area(){
	p := math.Phi
	fmt.Println( p * float64(c.r * 2))
}

func (s square) area(){
	fmt.Println(s.l * s.w)
}

func info(s shape){
	s.area()
}

func main() {
	s := square{
		l:10,
		w:20}
	c := circle{
		r:10,
	}
	//s.area()
	info(s)
	//c.area()
	info(c)
}
