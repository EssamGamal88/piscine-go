package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node.Left == nil {
		root = BTreeTransplant(root, node, node.Right)
	} else if node.Right == nil {
		root = BTreeTransplant(root, node, node.Left)
	} else {
		min := BTreeMin(node.Right)
		if min.Parent != node {
			root = BTreeTransplant(root, min, min.Right)
			min.Right = node.Right
			if min.Right != nil {
				min.Right.Parent = min
			}
		}
		root = BTreeTransplant(root, node, min)
		min.Left = node.Left
		if min.Left != nil {
			min.Left.Parent = min
		}
	}
	return root
}
