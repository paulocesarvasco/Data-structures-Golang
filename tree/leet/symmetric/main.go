package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return helper(root.Left, root.Right)
}

func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else if left.Val != right.Val {
		return false
	}
	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}
