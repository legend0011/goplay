package main

import (
	"fmt"
	"time"
)

func chanDeadlock() {
	// var c chan int  // c = nil
	c := make(chan int)
	// receive and send can't be in the same goroutine
	c <- 1
	<-c
}

func chanSimple() {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second)
		n := <-ch
		fmt.Printf("received: %d\n", n)
	}()
	ch <- 1 // must have launch receiver goroutine first?
	time.Sleep(time.Second)
}

func chanAsParam() {
	c := make(chan int)
	go func(ch chan int) {
		n := 1
		fmt.Printf("send: %d\n", n)
		ch <- n
	}(c)

	time.Sleep(time.Second)

	go func(ch chan int) {
		n := <-ch
		fmt.Printf("received: %d\n", n)
	}(c)

	time.Sleep(time.Second)
}
func main() {
	//chanDeadlock()
	//chanSimple()
	//chanAsParam()
	
}
