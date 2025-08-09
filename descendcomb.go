package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for a := 99; a >= 0; a-- {
		for b := a - 1; b >= 0; b-- {
			printTwoDigits(a)
			z01.PrintRune(' ')
			printTwoDigits(b)

			if !(a == 1 && b == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}

func printTwoDigits(n int) {
	z01.PrintRune(rune(n/10 + '0'))
	z01.PrintRune(rune(n%10 + '0'))
}
