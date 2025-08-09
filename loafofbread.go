package piscine

func LoafOfBread(str string) string {
	cleanStr := ""
	for _, r := range str {
		if r != ' ' {
			cleanStr += string(r)
		}
	}

	if len(cleanStr) == 0 {
		return "\n"
	}

	if len(cleanStr) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	charCount := 0

	for i := 0; i < len(str); i++ {
		char := str[i]

		if char == ' ' {
			continue
		}

		result += string(char)
		charCount++

		if charCount%5 == 0 {
			i++

			hasNextChar := false
			for j := i + 1; j < len(str); j++ {
				if str[j] != ' ' {
					hasNextChar = true
					break
				}
			}

			if hasNextChar {
				result += " "
			}
		}
	}

	return result + "\n"
}
