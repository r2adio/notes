package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func Main() {
	var count int // variable to hold the no. of lines to read from a file
	flag.IntVar(&count, "n", 5, "Numbers of line to read from a file")
	flag.Parse()

	var in io.Reader
	if filename := flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file: ", err)
		}
		defer f.Close()

		in = f
	} else {
		in = os.Stdin
	}

	buf := bufio.NewScanner(in)
	for range count {
		if !buf.Scan() {
			break
		}
		fmt.Println(buf.Text())
	}

	if err := buf.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading, err: ", err)
	}
}
