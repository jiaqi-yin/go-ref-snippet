package main

import (
	"context"
	"fmt"
	"time"
)

func req(ctx context.Context) {
	fmt.Println("Run goroutine")
	select {
	case <-ctx.Done():
		fmt.Println("Received cancel signal")
		return
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go req(ctx)

	num := 1
	for {
		if num > 3 {
			cancel()
		}
		time.Sleep(time.Second)
		num++
	}
}
