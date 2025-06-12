// NOTE: Error Package in go std lib:
// - check godoc for errors.New()
// - errors.New(): allows to create new error w/ just a string
// - also means, no need to define a new struct or type and have it explicitly implement the error interface
package main

import "errors"

// example for errors.New() fn
var err error = errors.New("something went wrong")
