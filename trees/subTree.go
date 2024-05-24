package trees

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root != nil && subRoot != nil {
		if root.Val == subRoot.Val {
			if isSameTree(root, subRoot) {
				return true
			}
		}
		return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
	}

	return false
}
