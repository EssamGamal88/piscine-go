package piscine

func ListForEach(l *List, f func(*NodeL)) {
	for node := l.Head; node != nil; node = node.Next {
		f(node)
	}
}

func Add2_node(node *NodeL) {
	switch v := node.Data.(type) {
	case int:
		node.Data = v + 2
	case string:
		node.Data = v + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch v := node.Data.(type) {
	case int:
		node.Data = v - 3
	case string:
		node.Data = v + "-3"
	}
}
