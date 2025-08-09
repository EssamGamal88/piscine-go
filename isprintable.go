package piscine

func IsPrintable(s string) bool {
	for _, ch := range s {
		if ch > '~' || ch < ' ' {
			return false
		}
	}
	return true
}
