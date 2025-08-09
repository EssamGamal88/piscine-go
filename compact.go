package piscine

func Compact(ptr *[]string) int {
	original := *ptr
	result := []string{}

	for _, str := range original {
		if str != "" {
			result = append(result, str)
		}
	}

	*ptr = result

	return len(result)
}
