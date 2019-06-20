package main

import "fmt"

type vehicle struct {
	doors string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	myTruck := truck{
		fourWheel: true,
		vehicle: vehicle{
				doors: "098",
				color: "dsdadsd",
			},
	}
	mySedan := sedan{
		luxury: true,
		vehicle: vehicle{
			doors: "2323",
			color: "c dds",
		},
	}
	fmt.Println(mySedan, myTruck)
	fmt.Println("Doors: ", mySedan.doors, " Color: ", mySedan.color, " Luxury: ", mySedan.luxury)
	fmt.Println("Doors: ", myTruck.doors, " Color: ", myTruck.color, " FourWheel: ", myTruck.fourWheel)
}
