package piscine

func BTreeIsBinary(root *TreeNode) bool {
	var isBinary func(root *TreeNode, min, max string) bool
	isBinary = func(root *TreeNode, min, max string) bool {
		if root == nil {
			return true
		}
		if root.Data <= min || root.Data >= max {
			return false
		}
		return isBinary(root.Left, min, root.Data) && isBinary(root.Right, root.Data, max)
	}

	return isBinary(root, "", string(rune(255)))
}
