// NOTE: structs are used to structure json data
package main

import (
	"fmt"
)

type rect struct {
	width  int
	height int
}

func (r *rect) area() int { return r.width * r.height }

func main() {
	// definition of structs in go
	type wheel struct {
		radius   int
		material string
	}

	// struct definition, w/ nested structs
	type car struct {
		brand      string
		model      string
		door       int
		mileage    int
		frontWheel wheel // nested structs
		backWheel  wheel // nested structs
	}

	// instantiate a new instance of a struct
	myCar := car{}
	myCar.frontWheel.radius = 5

	// anonymous structs: don't need a name, and imediately instantiating a new instance of a struct called mycar
	mycar := struct {
		brand string
		model string
	}{brand: "toyota", model: "camry"}
	fmt.Print(mycar.brand, mycar.model)

	// embeded structs: provide kinda data-only inheritance
	type CAR struct {
		make  string
		model string
	}
	type TRUCK struct {
		// "CAR" is embeded, so definition of TRUCK also contains all the fields of CAR struct
		CAR
		bedSize int
	}
	lanesTRUCK := TRUCK{
		bedSize: 10,
		CAR: CAR{
			make:  "toyota",
			model: "camry",
		},
	}
	fmt.Println(lanesTRUCK.bedSize)
	// embeded fields promted to top level
	// instead of lanesTRUCK.CAR.make
	fmt.Println(lanesTRUCK.make)
	fmt.Println(lanesTRUCK.model)

	// struct meathods
	// meathod declaration must be at package level, func (r *rect) area() int{}
	// type rect struct {
	// 	width int
	// 	height int
	// }
	// func (r *rect) area() int{
	// 	return r.width* r.height;
	// }
	r := rect{width: 5, height: 10}
	fmt.Println(r.area())
}
