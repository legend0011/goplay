package main

import (
	"fmt"
	"time"
)

// func chanDeadlock() {
// 	// var c chan int  // c = nil
// 	c := make(chan int)
// 	c <- 1 // here is the deadlock
// 	<-c
// }

// func chanSimple() {
// 	ch := make(chan int)
// 	go func() {
// 		for {
// 			n := <-ch
// 			fmt.Printf("received: %d\n", n)
// 		}
// 	}()
// 	ch <- 1 // must launch receiver goroutine first
// 	ch <- 2
// 	ch <- 3
// 	time.Sleep(time.Second) // otherwise, main thread don't have chance to run goroutine and exit
// }

type workerAndItem struct {
	w    int
	item int
}

func worker(workerid int, ch chan int, arr *[]workerAndItem) {
	for {
		n := <-ch
		*arr = append(*arr, workerAndItem{workerid, n})
		fmt.Printf("{%v} worker[%d] received: %d\n", time.Now().Nanosecond(), workerid, n)
		break
	}
}

func chanAsParam() {
	c := make(chan int)
	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			//fmt.Printf("send: %d\n", i)
			ch <- i
		}
	}(c)

	time.Sleep(time.Second)

	arr := make([]workerAndItem, 0)
	for i := 0; i < 1; i++ {
		go worker(i, c, &arr)
	}

	time.Sleep(time.Second)
	fmt.Println(arr)
}

func main() {
	//chanDeadlock()
	// chanSimple()
	chanAsParam()

}
