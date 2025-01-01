package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context timed out:", ctx.Err())
			return
		case t := <-ticker.C:
			fmt.Println("Ticker ticked at:", t)
		case <-timer.C:
			fmt.Println("Timer event triggered!")
			timer.Reset(5 * time.Second)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
