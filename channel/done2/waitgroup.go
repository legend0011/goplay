package main

import (
	"fmt"
	"sync"
)

func doWork(workerid int, ch <-chan int, wg *sync.WaitGroup) {
	for n := range ch {
		fmt.Printf("worker %d received: %d\n", workerid, n)
		wg.Done()
	}
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func CreateWorker(workerid int, wg *sync.WaitGroup) *worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	// receiver is in a goroutine
	go doWork(workerid, w.in, w.wg)
	return &w
}

func chanDemoWaitGroup() {
	var wg sync.WaitGroup

	var workers [10]*worker
	for i := 0; i < 10; i++ {
		workers[i] = CreateWorker(i, &wg)
	}

	// send job
	wg.Add(20)
	for i, worker := range workers {
		// send job
		worker.in <- i
	}
	for i, worker := range workers {
		// send job
		worker.in <- i + 20
	}

	// wait jobs to be done
	wg.Wait()
}

func main() {
	chanDemoWaitGroup()
}
