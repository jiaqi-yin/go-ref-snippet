package main

import (
	"context"
	"fmt"
	"time"
)

func req(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Timeout or cancel", ctx.Err().Error())
		return
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	go req(ctx)

	num := 1
	for {
		if num > 5 {
			cancel()
		}
		time.Sleep(time.Second)
		num++
	}
}
