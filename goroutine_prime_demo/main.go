package main

import "fmt"

func putNum(intChan chan<- int) {
	for i := 1; i < 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

// Test
func primeNum(intChan <-chan int, primeChan chan<- int, exitChan chan<- bool) {
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		isPrime := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primeChan <- num
		}
	}

	fmt.Println("primeNum goroutine exit")
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 4)

	go putNum(intChan)

	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("Prime=%d\n", res)
	}

	fmt.Println("Exit")
}
