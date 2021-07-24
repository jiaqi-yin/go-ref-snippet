package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int
	tail    int
}

func (queue *CircleQueue) Push(val int) error {
	if queue.IsFull() {
		return errors.New("queue is full")
	}
	queue.array[queue.tail] = val
	queue.tail = (queue.tail + 1) % queue.maxSize
	return nil
}

func (queue *CircleQueue) Pop() (int, error) {
	if queue.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	val := queue.array[queue.head]
	queue.head = (queue.head + 1) % queue.maxSize
	return val, nil
}

func (queue *CircleQueue) ListQueue() {
	size := queue.Size()
	if size == 0 {
		fmt.Println("queue is empty")
	}

	tempHead := queue.head
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d]=%d\t", tempHead, queue.array[tempHead])
		tempHead = (tempHead + 1) % queue.maxSize
	}
	fmt.Println()
}

func (queue *CircleQueue) IsFull() bool {
	return (queue.tail+1)%queue.maxSize == queue.head
}

func (queue *CircleQueue) IsEmpty() bool {
	return queue.tail == queue.head
}

func (queue *CircleQueue) Size() int {
	return (queue.tail + queue.maxSize - queue.head) % queue.maxSize
}

func main() {
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	var key string
	var val int
	for {
		fmt.Println("1. push to queue")
		fmt.Println("2. pop from queue")
		fmt.Println("3. list the queue")
		fmt.Println("4. exit")

		fmt.Scanln(&key)
		switch key {
		case "push":
			fmt.Println("enter an integer number:")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("added to the queue")
			}
		case "pop":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("get %d from the queue\n", val)
			}
		case "list":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
