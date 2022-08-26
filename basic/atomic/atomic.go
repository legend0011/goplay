package main

import (
	"fmt"
	"sync"
)

type Int int

func (a *Int) increment() {
	*a++ // read & write
}

func (a *Int) get() int {
	return int(*a)
}

//  go run -race basic/atomic/atomic.go
// check if there's data race.
func dataRace() {
	n := Int(0)
	threadTotal := 10
	wg := sync.WaitGroup{}
	wg.Add(threadTotal)
	for j := 0; j < threadTotal; j++ {
		go func() {
			for i := 0; i < 100; i++ {
				n.increment()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(n.get()) // 1000
}

type atomicInt struct {
	val  int
	lock sync.Mutex
}

func (a *atomicInt) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.val++
}

func (a *atomicInt) protectDataBlockWithLock() {
	fmt.Println("safe code block")
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.val++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.val
}

func noDataRace() {
	var n atomicInt
	total := 10
	wg := sync.WaitGroup{}
	wg.Add(total)
	for j := 0; j < total; j++ {
		go func() {
			for i := 0; i < 100; i++ {
				n.increment()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(n.get()) // 1000
}

func main() {
	//dataRace()
	noDataRace()
}
