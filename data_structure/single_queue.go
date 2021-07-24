package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array   [5]int
	front   int
	rear    int
}

func (queue *Queue) AddQueue(val int) error {
	if queue.rear == queue.maxSize-1 {
		return errors.New("queue is full")
	}

	queue.rear++
	queue.array[queue.rear] = val
	return nil
}

func (queue *Queue) GetQueue() (int, error) {
	if queue.front == queue.rear {
		return -1, errors.New("queue is empty")
	}
	queue.front++
	return queue.array[queue.front], nil
}

func (queue *Queue) ShowQueue() {
	fmt.Println("queue items...")
	for i := queue.front + 1; i <= queue.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, queue.array[i])
	}
	fmt.Println("")
}

func main() {
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}

	var key string
	var val int
	for {
		fmt.Println("1. add to queue")
		fmt.Println("2. get from queue")
		fmt.Println("3. show the queue")
		fmt.Println("4. exit")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("enter an integer number:")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("added to the queue")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("get %d from the queue\n", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
