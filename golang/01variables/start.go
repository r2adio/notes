package main

import "fmt"

func start() {
	//std out on console
	fmt.Println("startin go journey!")

	// basic variables types: bool int string float64 byte
	// declaring variable:
	// (sad variable declaration)
	var expLevel int // go gives value of 0, !garbage-value
	expLevel = 1
	// (modern variable declaration)
	futureExpLevel := 100 // :=, walrus operator
	pi := 3.14159
	message := "Hello, world!"
	isGoat := true
	fmt.Println(message, futureExpLevel, expLevel)
	fmt.Printf("value of pi= %.3f\n", pi)
	fmt.Printf("status: %t\n", isGoat)

	// go allows declaring multiple variables in a single line
	// mileage, company := 3434, "Toyota"

	// golang memory footprint
	// Go runtime: small amount of extra code that's included in the executable binary

	// constants in golang
	// constants cant use short declaration syntax
	const Pi = 3.145
}
