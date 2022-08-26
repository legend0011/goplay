package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator(seed int) chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i * seed
			i++
		}
	}()
	return out
}

func main() {
	ch1 := generator(1)
	ch2 := generator(2)
	timer := time.After(10 * time.Second) // will send to timer a signal after 10 sec
	for {
		select { // select who is faster
		case n := <-ch1:
			fmt.Printf("received: %d\n", n)
		case m := <-ch2:
			fmt.Printf("received: %d\n", m)
		case <-timer:
			fmt.Println("bye")
			return
		}
	}
}
