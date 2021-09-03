package main

import (
	"context"
	"fmt"
)

func req1(ctx context.Context) {
	fmt.Println(ctx.Value("start"))
	fmt.Println("Hello world.")

	ctx1 := context.WithValue(context.Background(), "next", "Finish context")
	req2(ctx1)
}

func req2(ctx context.Context) {
	fmt.Println(ctx.Value("next"))
}

func main() {
	ctx := context.WithValue(context.Background(), "start", "Golang context")
	req1(ctx)
}
