package main

import "fmt"

func if_else() {
	// if and else if statements
	var height int
	fmt.Print("enter the height: ")
	fmt.Scan(&height)
	if height > 6 {
		fmt.Println("you are super tall")
	} else if height > 4 {
		fmt.Println("you are tall enough")
	} else if height == 5 {
		fmt.Println("you've a perfect height")
	} else {
		fmt.Println("you are not tall enough")
	}
}
