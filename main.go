package main

import (
	"bufio"
	"fmt"
	"io"
	// "hello/queue"
	// "hello/tree"
)

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

type myReader struct {
	ComplexContent string
	Offset         int
	LenOnce        int
}

func (mr myReader) Read(p []byte) (n int, err error) {
	totalLen := len(mr.ComplexContent)
	if mr.Offset < totalLen {
		if leftLen := totalLen - mr.Offset; leftLen < mr.LenOnce {
			n = copy(p, []byte(mr.ComplexContent)[mr.Offset:mr.Offset+leftLen])
			return n, io.EOF
		}
		n = copy(p, []byte(mr.ComplexContent)[mr.Offset:mr.Offset+mr.LenOnce])
		return n, nil
	}
	return 0, nil
}

func printReaderContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	mr := myReader{ComplexContent: `jken
	xxx....
	emoji
	ha`,
		Offset:  0,
		LenOnce: 1,
	}

	printReaderContents(mr)

}
