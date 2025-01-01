package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type Incrementor struct {
	number int64
}

func NewIncrementor() *Incrementor {
	return &Incrementor{}
}

func (i *Incrementor) Add(value int64) {
	atomic.AddInt64(&i.number, value)
}

func (i *Incrementor) GetNumber() int64 {
	return atomic.LoadInt64(&i.number)
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

			localSum := int64(0)
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
}
