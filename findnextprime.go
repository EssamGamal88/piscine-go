package piscine

func FindNextPrime(nb int) int {
	for {
		if IsPrime(nb) {
			return nb
		}
		nb++
	}
}
