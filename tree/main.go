package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (n *Node) Insert(data int) {
	if data < n.data {
		if n.left == nil {
			n.left = &Node{data: data}
		} else {
			n.left.Insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data}
		} else {
			n.right.Insert(data)
		}
	}
}

func (n *Node) Search(data int) bool {
	if data == n.data {
		return true
	}
	if data < n.data && n.left != nil {
		return n.left.Search(data)
	} else if data < n.data && n.left == nil {
		return false
	} else if n.right != nil {
		return n.right.Search(data)
	} else {
		return false
	}
}

func (n *Node) Delete(data int) *Node {
	if data < n.data {
		if n.left != nil {
			n.left = n.left.Delete(data)
		}
	} else if data > n.data {
		if n.right != nil {
			n.right = n.right.Delete(data)
		}
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}
		minRight := n.right.Min()
		n.data = minRight.data
		n.right = n.right.Delete(minRight.data)
	}
	return n
}

func (n *Node) Depth() int {
	if n == nil {
		return 0
	}
	leftDepth := n.left.Depth()
	rightDepth := n.right.Depth()

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func (n *Node) Min() *Node {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current
}

func (n *Node) PrintInOrder() {
	if n.left != nil {
		n.left.PrintInOrder()
	}
	fmt.Println(n.data)
	if n.right != nil {
		n.right.PrintInOrder()
	}
}

func (n *Node) PrintPreOrder() {
	fmt.Println(n.data)
	if n.left != nil {
		n.left.PrintPreOrder()
	}
	if n.right != nil {
		n.right.PrintPreOrder()
	}
}

func (n *Node) PrintPostOrder() {
	if n.left != nil {
		n.left.PrintPostOrder()
	}
	if n.right != nil {
		n.right.PrintPostOrder()
	}
	fmt.Println(n.data)
}

func main() {
	root := &Node{data: 5}
	root.Insert(3)
	root.Insert(7)
	root.Insert(2)
	root.Insert(4)
	// root.Insert(6)
	root.Insert(9)
	root.Insert(8)
	root.Insert(10)
	fmt.Println(root.Depth())
}
