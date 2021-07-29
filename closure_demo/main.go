package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(fileName string) string {
		if strings.HasSuffix(fileName, suffix) {
			return fileName
		}
		return fileName + "." + suffix
	}
}

func main() {
	f := makeSuffix("jpg")
	fmt.Println(f("hello.jpg"))
	fmt.Println(f("world"))
	fmt.Println(f("text.txt"))
}
