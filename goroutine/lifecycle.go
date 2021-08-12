package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	runTimes := 0

	var wg sync.WaitGroup
	wg.Add(1)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine done")
				return
			default:
				fmt.Printf("Goroutine running times: %d\n", runTimes)
				runTimes++
			}
			if runTimes > 5 {
				cancel()
				wg.Done()
			}
		}
	}(ctx)
	wg.Wait()
}
