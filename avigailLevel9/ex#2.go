package main

import (
	"fmt"
)

type person struct {
	name string
}

func (p *person) speak(){
	fmt.Println("speak ",p.name)
}

type human interface {
	speak()
}

func saySomething(h human){
	h.speak()
}
func main() {
	p1 := person{
		name: "Avigail",
	}
	saySomething(&p1)
	p1.speak()
	//saySomething(p1)


}
