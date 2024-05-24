package trees

func isSameTree(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 != nil && t2 != nil {
		return (t1.Val == t2.Val) && isSameTree(t1.Left, t2.Left) && isSameTree(t1.Right, t2.Right)
	}
	return false
}
