package piscine

func AlphaCount(s string) int {
	count := 0
	for _, i := range s {
		if (i >= 'A' && i <= 'Z') || (i >= 'a' && i <= 'z') {
			count++
		}
	}
	return count
}
