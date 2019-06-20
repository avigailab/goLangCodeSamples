package main

import "fmt"

func main() {
	myTruck := struct{
		luxury bool
		doors int
		color string
	}{luxury: true,
		doors: 1,
		color: "c dds",
	}
	fmt.Println(myTruck)
	fmt.Println("Doors: ", myTruck.doors, " Color: ", myTruck.color, " Luxury: ", myTruck.luxury)
}
