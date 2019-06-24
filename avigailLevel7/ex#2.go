package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func changeMe(p *person){
	p.first = fmt.Sprintf("%v - changed ", p.first)
	p.last = fmt.Sprintf("%v - changed ", p.last)
	p.age++
	//(*p).first = fmt.Sprintf("%v - changed ", p.first)
	//(*p).last = fmt.Sprintf("%v - changed ", p.last)
	//(*p).age++

}
func main() {
	p1 := person{
		first: "avigail",
		last: "a",
		age: 40,
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
