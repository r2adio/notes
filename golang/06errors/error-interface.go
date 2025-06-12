// NOTE: Custom Error Interface
// - 'cause errors are just Interface, we can build our own custom types that implements the `error` interface.
package main

import "fmt"

// eg: `userError` struct that implements`error` interface
type userError struct {
	name string
}

func (e userError) Error() string {
	return fmt.Sprintf("%v has a problem w/ their account", e.name)
}
