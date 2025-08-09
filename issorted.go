package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	isAsc := true
	isDesc := true

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			isAsc = false
		}
		if f(a[i], a[i+1]) < 0 {
			isDesc = false
		}
	}

	return isAsc || isDesc
}
