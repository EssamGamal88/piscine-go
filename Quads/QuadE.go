package main

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 && col == 0 || row == y-1 && col == x-1 && x > 1 && y > 1 {
				z01.PrintRune('A')
			} else if row == 0 && col == x-1 && x > 1 || row == y-1 && col == 0 {
				z01.PrintRune('C')
			} else if row == 0 || row == y-1 && (col == 1 || col == 2 || col == 3) || col == 0 || col == x-1 {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
