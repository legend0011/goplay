package main

import (
	"fmt"
	"time"
)

func asyncIO() {
	for i := 0; i < 10; i++ {
		go func(taskid int) {
			for {
				// printf is io operation, will always have chance to switch to other goroutine
				fmt.Printf("hello from %d\n", taskid)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}

// main thread is also a goroutine
//  i will reach 10 in main goroutine
//  i is also used in goroutine
func goroutineNoHangOverOverflow() {
	// difference between
	//   var a [10]int
	//   a := [10]int // this is wrong
	//   a := make([]int, 10)
	var a [10]int
	for i := 0; i < 10; i++ {
		go func() {
			for {
				// don't have chance to switch to other goroutine
				a[i]++
			}
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

func goroutineNoHangOver() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(idx int) {
			for {
				// don't have chance to switch to other goroutine
				a[idx]++
				//// manually switch
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

func main() {
	// asyncIO()

	goroutineNoHangOver()
}
