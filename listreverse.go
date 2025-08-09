package piscine

func ListReverse(l *List) {
	var prev *NodeL = nil
	current := l.Head
	var next *NodeL = nil
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Tail = l.Head
	l.Head = prev
}
