package main

import "fmt"

var version = "1.2"

func main() {
	fmt.Println("Hello world ", version)
}

// go build -o test main.go
// Output: Hello world  1.2

// go build -ldflags "-X main.version=1.2.3" -o test main.go
// Output: Hello world  1.2.3
