package main

import (
	"fmt"
)

func writeData(intChan chan int, size int) {
	for i := 0; i < size; i++ {
		fmt.Println("Write data: ", i)
		intChan <- i
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("Read data: ", v)
	}

	exitChan <- true
	close(exitChan)
}

func main() {
	size := 10
	exitChan := make(chan bool, 1)
	intChan := make(chan int, size)
	go writeData(intChan, size)
	go readData(intChan, exitChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
	fmt.Println("Exit")
}
