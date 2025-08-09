package piscine

func Compare(a, b string) int {
	astring := []rune(a)
	bstring := []rune(b)

	for i := 0; i < len(astring) && i < len(bstring); i++ {
		if astring[i] < bstring[i] {
			return -1
		} else if astring[i] > bstring[i] {
			return 1
		}
	}
	if len(astring) < len(bstring) {
		return -1
	} else if len(astring) > len(bstring) {
		return 1
	}
	return 0
}
