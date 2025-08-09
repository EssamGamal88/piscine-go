package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium)/2; i++ {
		j := len(podium) - 1 - i
		podium[i], podium[j] = podium[j], podium[i]
	}
	return podium
}
