package piscine

func TrimAtoi(s string) int {
	sign := 1
	number := 0
	flag := false

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			flag = true
			number = number*10 + int(ch-'0')
		} else if ch == '-' && !flag {
			sign = -1
		}
	}

	return number * sign
}
