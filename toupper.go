package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			runes[i] = ch - 32
		}
	}
	return string(runes)
}
