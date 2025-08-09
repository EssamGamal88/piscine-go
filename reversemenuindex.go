package piscine

func ReverseMenuIndex(menu []string) []string {
	length := len(menu)
	reversed := make([]string, length)

	for i := 0; i < length; i++ {
		reversed[i] = menu[length-1-i]
	}

	return reversed
}
