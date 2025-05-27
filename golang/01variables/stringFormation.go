package main

import "fmt"

func stringFormatting() {
	const name = "Saul Goodman"
	const openRate = 30.5

	// returns the formatted string
	msg := fmt.Sprintf("Hi %v, your open rate is %.1f percent\n", name, openRate)

	// prints a formatted string to standard out
	fmt.Print(msg)
}
