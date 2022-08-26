package main

import (
	"fmt"
	"sync"
	"time"
)

var n = 0
var once sync.Once

func init0() {

	once.Do(func() {
		fmt.Printf("initializing %d time!\n", n)
		n += 1
	})
}

func init1() {
	fmt.Printf("initializing %d time!\n", n)
	n += 1
}

func main() {
	for i := 0; i <= 10; i++ {
		//go init1()
		go init0()
	}
	time.Sleep(time.Second * 3)
}
