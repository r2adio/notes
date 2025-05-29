// NOTE: both the functions are same
package main

func named_return_values() (x, y int) {
	// x and y are initialized w/ 0 values
	return // automatically returns x and y
}
func return_values() (int, int) {
	var x int
	var y int
	return x, y // using explicit returns
}

// NOTE:
// - return statements w/o arguments returns the named return values, called "naked return".
// - naked return statements should be used only in  smaller functions, as they harm readability

// NOTE: Benefits of named returns
// - good for understading(documentation)
//   - func calculator(a, b int) (mul, div int, err error) {... return mul,div, nil}
//   - func calculator(a, b int) (int, int, error) {... return mul, div, nil}
//	 - using named return parameters,first, are particularly impt for longer func, which is easier to understad than latter
//	 - using named return parameters, meaning of each return value is clear just by looking at function signature
