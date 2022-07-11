package main

import "hello/queue"

func main() {
	q := queue.Queue{1, 2}
	q.Push(3)
	q.Print()
}
