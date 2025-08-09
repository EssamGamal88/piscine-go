package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	newWord := true

	for i, ch := range runes {
		if isAlphaNumeric(ch) {
			if newWord {
				if ch >= 'a' && ch <= 'z' {
					runes[i] = ch - 32
				}
				newWord = false
			} else {
				if ch >= 'A' && ch <= 'Z' {
					runes[i] = ch + 32
				}
			}
		} else {
			newWord = true
		}
	}
	return string(runes)
}

func isAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9')
}
