// NOTE: switch case
// - break statement isn't required, the break statement is implicit in go
// - dont want a case to fall-through to the next case, use 'fallthrough' keyword
package main

func switch_case(os string) string {
	var creator string
	switch os {
	case "linux":
		creator = "Linus Torvalds"
	case "windows":
		creator = "Bill Gates"
	case "mac":
		creator = "A Steve"
	default:
		creator = "unknown"
	}
	return creator
}
