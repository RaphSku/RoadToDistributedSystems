package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Incrementor struct {
	mtx    sync.Mutex
	number int
}

func NewIncrementor() *Incrementor {
	return &Incrementor{}
}

func (i *Incrementor) Add(value int) {
	i.mtx.Lock()
	defer i.mtx.Unlock()
	i.number += value
}

func (i *Incrementor) GetNumber() int {
	i.mtx.Lock()
	defer i.mtx.Unlock()
	return i.number
}

func main() {
	target := 1_000_000
	maxProcs := runtime.GOMAXPROCS(0)

	inc := NewIncrementor()
	workPerGoroutine := target / maxProcs
	remainder := target % maxProcs

	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(maxProcs)
	for i := 0; i < maxProcs; i++ {
		start := i * workPerGoroutine
		end := start + workPerGoroutine
		if i == maxProcs-1 {
			end += remainder
		}

		go func(start, end int) {
			defer wg.Done()

			localSum := 0
			for j := start; j < end; j++ {
				localSum += 1
			}

			inc.Add(localSum)
		}(start, end)
	}
	wg.Wait()
	elapsed := time.Since(start)

	fmt.Printf("Elapsed time (parallel): %s\n", elapsed)
	fmt.Printf("Number reached (parallel): %d\n", inc.GetNumber())

	// HERE WE SIMPLY INCREMENT AND COMPARE IT
	start = time.Now()
	sum := 0
	for i := 0; i < target; i++ {
		sum += 1
	}
	elapsed = time.Since(start)

	fmt.Printf("Elapsed time (sequentially): %s\n", elapsed)
	fmt.Printf("Number reached (sequentially): %d\n", sum)
}
