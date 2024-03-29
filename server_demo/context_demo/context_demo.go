package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("done!") // first print this to show goroutine died, then main() done
					return               // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	func() {
		defer cancel() // cancel when we are finished consuming integers

		for n := range gen(ctx) {
			fmt.Println(n)
			if n == 5 {
				break
			}
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("main() done!")
}
