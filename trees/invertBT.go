package trees

func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	temp := invertTree(root.Left)
	root.Left = invertTree(root.Right)
	root.Right = temp

	return root
}
