package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func chanOfInt() {
	quit := make(chan int)

	for i := 0; i < 200; i++ {
		go func(ith int) {
			select {
			case <-quit:
				fmt.Printf("[%d] quit\n", ith)
			}
		}(i)
	}

	timeoutChan := time.After(5 * time.Second)
	select {
	case <-timeoutChan:
		quit <- 2
	}

	time.Sleep(time.Second) // wait go routine to finish execution
	fmt.Println("over!")
}

func closeChanOfInt() {
	quit := make(chan int)

	for i := 0; i < 10; i++ {
		go func(ith int) {
			select {
			case <-quit:
				fmt.Printf("[%d] quit\n", ith)
			}
		}(i)
	}

	timeoutChan := time.After(5 * time.Second)
	select {
	case <-timeoutChan:
		close(quit)
	}

	time.Sleep(time.Second) // wait go routine to finish execution
	fmt.Println("over!")
}

func closeChanOfEmpty() {
	quit := make(chan struct{}) // only care about close signal

	for i := 0; i < 10; i++ {
		go func(ith int) {
			select {
			case <-quit:
				fmt.Printf("[%d] quit\n", ith)
			}
		}(i)
	}

	timeoutChan := time.After(5 * time.Second)
	select {
	case <-timeoutChan:
		close(quit)
	}

	time.Sleep(time.Second) // wait go routine to finish execution
	fmt.Println("over!")
}

func closeChanOfEmptyWithWaitGroup() {
	quit := make(chan struct{})

	wg := sync.WaitGroup{}
	numOfGoR := 5
	for i := 0; i < numOfGoR; i++ {
		wg.Add(1)
		go func(ith int) {
			select {
			case <-quit:
				fmt.Printf("[%d] quit!\n", ith)
			}
			wg.Done()
		}(i)
	}

	// send quit signal
	close(quit)

	wg.Wait()
	fmt.Println("over!!")
}

func contextControlTimeoutDemo() {
	// empty parent ctx
	pctx := context.TODO() // or can use context.Background()

	ctx, _ := context.WithTimeout(pctx, 5*time.Second)

	numOfGoR := 5
	for i := 0; i < numOfGoR; i++ {
		go func(ith int) {
			select {
			case <-ctx.Done():
				fmt.Printf("[%d] quit!\n", ith)
			}
		}(i)
	}

	// send quit signal
	time.Sleep(6 * time.Second)
	fmt.Println("over!!")
}

func contextControlCancelDemo() {
	// empty parent ctx
	pctx := context.TODO() // or can use context.Background()

	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)

	numOfGoR := 5
	for i := 0; i < numOfGoR; i++ {
		go func(ith int) {
			select {
			case <-ctx.Done():
				fmt.Printf("[%d] quit!\n", ith)
			}
		}(i)
	}
	// In reallife, if error happens,
	//  cancel() will cancel the go routine immediately, instead of 5 seconds timeout
	cancel()

	// send quit signal
	time.Sleep(6 * time.Second)
	fmt.Println("over!!")
}

func contextDemoWithCancel() {
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
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func contextControlWithTimeoutDemo() {
	exceedDeadline := func(ctx context.Context) bool {
		select {
		case <-ctx.Done():
			fmt.Printf("ctx.Err() = %s\n", ctx.Err())
			return true
		default: // The default case in a select is run if no other case is ready.
			fmt.Println("fine")
			return false
		}
	}

	heavyWork := func() {
		time.Sleep(5 * time.Second)
		fmt.Println("already sleep 5...")
	}

	do := func(ctx context.Context) {
		// check if exceed deadline during each heavyWork().
		// if so, return to release goroutine resources
		heavyWork()
		if exceedDeadline(ctx) {
			return
		}

		heavyWork()
		if exceedDeadline(ctx) {
			return
		}

		heavyWork()
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 8*time.Second)
	defer cancel()
	//cancel()
	go do(ctx)
	time.Sleep(50 * time.Second)
}

func contextWithValueDemo() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))
}

func main() {
	// read this first: https://pkg.go.dev/context#pkg-overview
	// Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context.
	// The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue.
	// When a Context is canceled, all Contexts derived from it are also canceled.

	//chanOfInt()
	//closeChanOfInt()
	//closeChanOfEmpty() // only care about close signal
	//closeChanOfEmptyWithWaitGroup()
	//contextControlTimeoutDemo()
	//contextControlCancelDemo()
	contextControlWithTimeoutDemo()
	//contextWithValueDemo()
}
