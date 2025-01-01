package main

import (
	"fmt"
	"sync"
)

type valueCommunicator struct {
	communicationCh chan uint8
}

func newValueCommunicator() *valueCommunicator {
	return &valueCommunicator{
		communicationCh: make(chan uint8),
	}
}

func (vc *valueCommunicator) valueGenerator(value uint8) {
	vc.communicationCh <- value % 10
}

func (vc *valueCommunicator) valueConsumer() {
	value := <-vc.communicationCh
	fmt.Printf("Value received: %d\n", value)
}

func main() {
	values := []uint8{0, 10, 20, 25, 50, 66, 71, 96, 124, 127, 255}

	vc := newValueCommunicator()
	for _, value := range values {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			vc.valueGenerator(value)
		}()
		go func() {
			defer wg.Done()
			vc.valueConsumer()
		}()
		wg.Wait()
	}
}
