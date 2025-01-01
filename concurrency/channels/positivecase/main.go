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

func (vc *valueCommunicator) valueGenerator(values []uint8) {
	for _, value := range values {
		vc.communicationCh <- value % 10
	}
	close(vc.communicationCh)
}

func (vc *valueCommunicator) valueConsumer() {
	for value := range vc.communicationCh {
		fmt.Printf("Value received: %d\n", value)
	}
}

func main() {
	values := []uint8{0, 10, 20, 25, 50, 66, 71, 96, 124, 127, 255}

	vc := newValueCommunicator()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		vc.valueGenerator(values)
	}()
	go func() {
		defer wg.Done()
		vc.valueConsumer()
	}()
	wg.Wait()
}
