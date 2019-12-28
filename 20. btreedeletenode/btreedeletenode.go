package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			current := root.Right
			root = nil
			return current
		} else if root.Right == nil {
			current := root.Left
			root = nil
			return current
		}
		current := BTreeMin(root.Right)
		root.Data = current.Data
		root.Right = BTreeDeleteNode(root.Right, current)
	}
	return root
}
