package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	lastSlash := 0

	for i, c := range name {
		if c == '/' {
			lastSlash = i + 1
		}
	}

	for _, c := range name[lastSlash:] {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
