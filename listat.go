package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	current := l
	index := 0

	for current != nil {
		if index == pos {
			return current
		}
		current = current.Next
		index++
	}
	return nil
}
