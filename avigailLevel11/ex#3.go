package main

import (
	"errors"
	"fmt"
	)

type customErr struct {
	err error
	data string
}

func (ce customErr) Error() string {
	return fmt.Sprintln("Error Occurred", ce.err, ce.data)
}

func main() {
	c1 := customErr{
		err: errors.New("err"),
		data: "my error data"}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
