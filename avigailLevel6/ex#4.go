package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

func (p person) speak(){
	fmt.Printf("%v %v says %d", p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "avigaila",
		last: "abuhai",
		age: 26}
	p1.speak()
}
