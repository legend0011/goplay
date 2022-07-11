// package main

// import (
// 	"fmt"
// 	"hello/queue"
// 	"hello/tree"
// )

// type MyTreeNode struct {
// 	Node *tree.TreeNode
// }

// func (mynode *MyTreeNode) PostOrder() {
// 	if mynode == nil || mynode.Node == nil {
// 		return
// 	}
// 	left := MyTreeNode{Node: mynode.Node.Left}
// 	right := MyTreeNode{Node: mynode.Node.Right}
// 	left.PostOrder()
// 	right.PostOrder()
// 	fmt.Println(mynode.Node.Val)
// }

// func main() {
// 	root := tree.TreeNode{Val: 1}
// 	root.Left = &tree.TreeNode{Val: 2}
// 	root.Right = &tree.TreeNode{Val: 5}
// 	root.Traverse()

// 	myRoot := &MyTreeNode{Node: &root}
// 	myRoot.PostOrder()

// 	q := queue.Queue{1}
// }
