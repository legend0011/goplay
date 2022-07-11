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
