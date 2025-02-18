package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) Insert(val int) {
	if val < n.Val {
		if n.Left == nil {
			n.Left = &TreeNode{Val: val}
		} else {
			n.Left.Insert(val)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Val: val}
		} else {
			n.Right.Insert(val)
		}
	}
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	p := &TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}

	q := &TreeNode{Val: 1}
	q.Insert(2)

	fmt.Println(isSameTree(p, q))
}
