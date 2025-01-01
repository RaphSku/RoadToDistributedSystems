package main

import (
	"fmt"
	"os"
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

func averageDurationTimes(durations []time.Duration) time.Duration {
	if len(durations) == 0 {
		fmt.Printf("Error computing the average over times since times is empty slice!")
		os.Exit(1)
	}

	var total time.Duration
	for _, d := range durations {
		total += d
	}

	return total / time.Duration(len(durations))
}

func main() {
	targets := []int{1, 10, 50, 100, 250, 500, 750, 1_000, 5_000, 10_000, 50_000, 100_000, 500_000, 1_000_000, 5_000_000, 10_000_000, 50_000_000, 100_000_000}
	samplesToAverage := 100
	maxProcs := runtime.GOMAXPROCS(0)

	file, err := os.Create("benchmark.log")
	if err != nil {
		fmt.Printf("Error during creation of benchmark file!")
		os.Exit(1)
	}
	defer file.Close()

	content := fmt.Sprintf("Number of samples for each target: %d\n", samplesToAverage)
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Could not log to benchmark file!")
		os.Exit(1)
	}

	for _, target := range targets {
		timeMsParallel := []time.Duration{}
		timeMsSequential := []time.Duration{}
		for range samplesToAverage {
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
			timeMsParallel = append(timeMsParallel, elapsed)

			start = time.Now()
			sum := 0
			for i := 0; i < target; i++ {
				sum += 1
			}
			elapsed = time.Since(start)
			timeMsSequential = append(timeMsSequential, elapsed)
		}
		averagedParallelTimes := averageDurationTimes(timeMsParallel)
		content := fmt.Sprintf("Target: %d\nElapsed Time (parallel): %s\n", target, averagedParallelTimes)
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Printf("Could not log to benchmark file!")
			os.Exit(1)
		}

		averagedSequentialTimes := averageDurationTimes(timeMsSequential)
		content = fmt.Sprintf("Elapsed Time (sequential): %s\n", averagedSequentialTimes)
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Printf("Could not log to benchmark file!")
			os.Exit(1)
		}
	}
}
