// NOTE: Errors:
// - when something went wrong in a fn, that fn should return an `error` as its last return value
package main

func main() {
	// go express errors w/ `error` values
	// Error is any type that implements simple built-in `error interface`
	type error interface {
		Error() string
	}
}
