package main

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func main() {
	useColor := flag.Bool("color", false, "display colorized output")
	flag.Parse()

	if *useColor {
		colorize(ColorGreen, "Colored Text!")
		return
	}
	// using ANSI escape codes
	fmt.Println("\033[34mThis text is blue.\033[0m")

	// using third-party library
	color.Cyan("hell")
}
