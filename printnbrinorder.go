package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var digits [10]int
	for n > 0 {
		d := n % 10
		digits[d]++
		n /= 10
	}

	for i := 0; i < 10; i++ {
		for digits[i] > 0 {
			z01.PrintRune(rune(i) + '0')
			digits[i]--
		}
	}
}
