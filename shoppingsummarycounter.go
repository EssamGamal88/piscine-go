package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	counts := make(map[string]int)
	wordStarted := false
	word := ""

	for i := 0; i < len(str); i++ {
		char := str[i]
		if char != ' ' {
			word += string(char)
			wordStarted = true
		} else {
			if wordStarted {
				counts[word]++
				word = ""
				wordStarted = false
			}
		}
	}

	if wordStarted {
		counts[word]++
	}

	return counts
}
