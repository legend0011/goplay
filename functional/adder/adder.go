package main

import (
	"fmt"
)

// adder returns an environment: free variable + method
func adder() func(int) int {
	sum := 0
	fmt.Printf("run again..\n")
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("1 + ... + %d = %d\n", i, a(i))
	}

	fmt.Println("1 ", adder()(1))
	fmt.Println("1. ", adder()(1))
	fmt.Println("1.. ", adder()(1))

}
