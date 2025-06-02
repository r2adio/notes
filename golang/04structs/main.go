package main

func main() {
	// nested structs
	type wheel struct {
		radius   int
		material string
	}

	// definition of structs in go
	type car struct {
		brand      string
		model      string
		door       int
		mileage    int
		frontWheel wheel
		backWheel  wheel
	}

	myCar := car{}
	myCar.frontWheel.radius = 5

	mycar := struct {
		brand string
		model string
	}{brand: "toyota", model: "camry"}
}
