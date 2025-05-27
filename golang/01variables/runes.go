// NOTE:
// - in C, char is a single byte. usign ascii encoding, the standard for C, we can represent 128 characters
// - in go, string are sequence of bytes, they can hold arbitrary data.
// - runes, alias for int32, means rune is a 32-bit int, which is large enough to hold any unicode code point
// - usage of runes: when work with individual characters in a string
// - with help of runes, go can handle wide variety of unicode characters in strings
package main

import (
	"fmt"
	"unicode/utf8"
)

func runes() {
	// const name = "boots"
	const name = "🐻"
	fmt.Printf("constant 'name' byte length: %d\n", len(name))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name)
}
