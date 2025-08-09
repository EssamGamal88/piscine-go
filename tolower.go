package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for i, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			runes[i] = ch + 32
		}
	}
	return string(runes)
}
