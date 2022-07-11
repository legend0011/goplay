package queue

import "fmt"

type Queue []int

func (q *Queue) Push(val int) {
	*q = append(*q, val)
}

func (q *Queue) Print() {
	for i, v := range *q {
		fmt.Printf("(%d, %d) ", i, v)
	}
}
