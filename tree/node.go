package tree

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}

	fmt.Println(node.Val)
	node.Left.Traverse()
	node.Right.Traverse()
}

func (node *TreeNode) TraverseWithFunc(f func(*TreeNode)) {
	if node == nil {
		return
	}
	node.Left.TraverseWithFunc(f)
	f(node)
	node.Right.TraverseWithFunc(f)
}

func (node *TreeNode) TraverseWithChan() chan *TreeNode {
	out := make(chan *TreeNode)
	go func() {
		node.TraverseWithFunc(func(node *TreeNode) {
			out <- node
		})
		close(out)
	}()
	return out
}
