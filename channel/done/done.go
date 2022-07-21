package main

import (
	"fmt"
)

func doWork(workerid int, ch <-chan int, done chan<- bool) {
	for n := range ch {
		fmt.Printf("worker %d received: %d\n", workerid, n)
		// use communication to share memory, not the way around
		done <- true
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func CreateWorker(workerid int) *worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	// receiver is in a goroutine
	go doWork(workerid, w.in, w.done)
	return &w
}

func chanDemo() {
	var workers [10]*worker
	for i := 0; i < 10; i++ {
		workers[i] = CreateWorker(i)
	}

	for i, worker := range workers {
		// send job
		worker.in <- i
	}
	// must add this code block here, because
	//   for each worker, it has its own channel
	//   after worker[i] have done the job , worker[i] then send out done=true
	//   in the main goroutine, worker[i]'s done must be consumed to continue sending new jobs
	// ---------
	//  worker: 1  2  3  4
	//  job:    |  |  |  |
	//  done:   d  d  d  d
	// new job: |  |  |  |
	//  done:   d  d  d  d
	for _, worker := range workers {
		// wait to be done
		<-worker.done
	}

	for i, worker := range workers {
		// send job
		worker.in <- i + 100
	}
	for _, worker := range workers {
		// wait to be done
		<-worker.done
	}
}

func main() {
	chanDemo()
	// bufferedChannel()
}
