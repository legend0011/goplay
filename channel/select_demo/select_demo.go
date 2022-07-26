package main

import (
	"fmt"
)

func fib(ch chan<- int, quit <-chan int) {
	x, y := 0, 1
	for {
		select { // select can select producer and consumer chan operations
		case <-quit:
			fmt.Println("quit!")
			return
		case ch <- x:
			x, y = x+y, x
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan int)

	// consumer can create channel too
	// start consumer async-ly first
	go consume(ch, quit)

	fib(ch, quit)
}

func consume(ch <-chan int, quit chan<- int) {
	total := 10
	for i := range ch {
		if total == 0 {
			break
		}
		fmt.Println(i)
		total--
	}
	quit <- 0
}
