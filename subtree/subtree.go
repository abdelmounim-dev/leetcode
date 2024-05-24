package subtree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 != nil && t2 != nil {
		return (t1.Val == t2.Val) && isSameTree(t1.Left, t2.Left) && isSameTree(t1.Right, t2.Right)
	}
	return false
}

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
