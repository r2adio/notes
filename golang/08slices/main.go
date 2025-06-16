// NOTE: Slices syntax
// - arrayName[lowIndex:highIndex]
// - arrayName[lowIndex:]
// - arrayName[:highIndex]
// - arrayName[:]
package main

import "fmt"

func main() {
	// declaring array of 10 integers
	var qwe [10]int
	// initialize array of size 4
	abc := [4]int{1, 2, 3, 4}
	// using slices
	ABC := abc[1:3] // [2, 3, 4]
	ABC = abc[:]    // [1, 2, 3, 4]
	fmt.Println(qwe, ABC)

	// create a new slice w/o explicitly creating array under the hood
	mySlice := make([]int, 5, 10) // new slice of int, initial length of 5, and capacity of 10
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	// here capacity can be understood as length of the underlying array
	// slices created w/ `make` will be filled w/ zero value of the type

	// capacity argument is usually omitted and defaults to the length of slice
	myslice := make([]int, 5)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))

	// creating slice w/ a specific set of values
	sliceString := []string{"hi", "there", "!"} // using [3]string would make it array instead of slice
	fmt.Println(len(sliceString))               // 3
	fmt.Println(cap(sliceString))
}
