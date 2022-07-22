package main

import (
	"fmt"
	"hello/tree"
	"time"
)

func main() {
	root := tree.TreeNode{Val: 1}
	root.Left = &tree.TreeNode{Val: 2}
	root.Right = &tree.TreeNode{Val: 5}

	root.Traverse()

	sum := 0
	root.TraverseWithFunc(func(node *tree.TreeNode) {
		sum = sum + node.Val
	})
	fmt.Printf("total sum = %d of nodes' value\n", sum)

	fmt.Println("TraverseWithChan")
	sum = 0
	ch := root.TraverseWithChan()
	time.Sleep(time.Second)
	for node := range ch {
		sum = sum + node.Val
	}
	fmt.Printf("chan: total sum = %d of nodes' value\n", sum)
}
