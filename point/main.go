package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	ptr := &point{}
	setPoint(ptr)

	var chars [15]rune

	chars[0] = 120
	chars[1] = 32
	chars[2] = 61
	chars[3] = 32
	chars[4] = 52
	chars[5] = 50
	chars[6] = 44
	chars[7] = 32
	chars[8] = 121
	chars[9] = 32
	chars[10] = 61
	chars[11] = 32
	chars[12] = 50
	chars[13] = 49
	chars[14] = 10

	for i := 0; i < 15; i++ {
		z01.PrintRune(chars[i])
	}
}
