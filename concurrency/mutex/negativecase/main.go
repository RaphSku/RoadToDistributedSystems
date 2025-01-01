package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Incrementor struct {
	mtx    sync.Mutex
	number int
}

func NewIncrementor() *Incrementor {
	return &Incrementor{}
}

func (i *Incrementor) Increment() {
	i.mtx.Lock()
	defer i.mtx.Unlock()
	i.number += 1
}

func (i *Incrementor) GetNumber() int {
	i.mtx.Lock()
	defer i.mtx.Unlock()
	return i.number
}

func main() {
	maxProcs := runtime.GOMAXPROCS(0)
	fmt.Printf("Number of cores: %d\n", maxProcs)
	target := 1000

	totalIncrementCounter := 0
	var totalIncrementCounterMtx sync.Mutex

	inc := NewIncrementor()
	var wg sync.WaitGroup
	wg.Add(maxProcs)
	for i := 0; i < maxProcs; i++ {
		go func() {
			defer wg.Done()
			for {
				totalIncrementCounterMtx.Lock()
				if inc.GetNumber() >= target {
					totalIncrementCounterMtx.Unlock()
					break
				}
				totalIncrementCounter += 1
				totalIncrementCounterMtx.Unlock()

				inc.Increment()
			}
		}()
	}
	wg.Wait()

	fmt.Printf("Number reached: %d\n", inc.GetNumber())
}
