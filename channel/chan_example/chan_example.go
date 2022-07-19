package main

import (
	"fmt"
	"time"
)

func worker(workerid int, ch <-chan int) {
	// for {
	// 	n, ok := <-ch
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Printf("worker %d received: %c\n", workerid, n)
	// }
	for n := range ch {
		fmt.Printf("worker %d received: %c\n", workerid, n)
	}
}

func CreateWorker(workerid int) chan<- int {
	ch := make(chan int)
	// receiver is in a goroutine
	go worker(workerid, ch)
	return ch
}

// func chanDemo() {
// 	var channels [10]chan<- int
// 	for i := 0; i < 10; i++ {
// 		channels[i] = CreateWorker(i)
// 	}

// 	for i := 0; i < 10; i++ {
// 		channels[i] <- 'a' + i
// 	}

// 	for i := 0; i < 10; i++ {
// 		channels[i] <- 'A' + i
// 	}

// 	time.Sleep(time.Second)

// }

func bufferedChannel() {
	// improve efficiency
	ch := make(chan int, 3)
	go worker(0, ch)
	ch <- 'e'
	ch <- 'b'
	ch <- 'c'
	ch <- 'd'
	close(ch)
	time.Sleep(time.Millisecond)
}

func main() {
	// chanDemo()
	bufferedChannel()
}
