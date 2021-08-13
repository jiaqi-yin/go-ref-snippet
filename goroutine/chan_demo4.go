package main

import "fmt"

func main() {
	count := 5
	eggs := make(chan int, count)
	exit := make(chan bool, count)

	go func() {
		for i := 0; i < count; i++ {
			eggs <- i
		}
		close(eggs)
	}()

	for i := 0; i < 100; i++ {
		go func(num int) {
			if egg, ok := <-eggs; ok {
				fmt.Printf("%d got egg %d\n", num, egg)
				exit <- true
			}
		}(i)
	}

	for i := 0; i < count; i++ {
		<-exit
	}

	fmt.Println("Exit")

}
