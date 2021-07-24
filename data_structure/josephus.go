package main

import "fmt"

type Node struct {
	Id   int
	Next *Node
}

func AddNode(num int) *Node {
	first := &Node{}
	current := &Node{}

	if num < 1 {
		fmt.Println("num should be great than zero")
		return first
	}

	for i := 1; i <= num; i++ {
		newNode := &Node{Id: i}
		if i == 1 {
			first = newNode
			current = newNode
			current.Next = first
		} else {
			current.Next = newNode
			current = newNode
			current.Next = first
		}
	}
	return first
}

func ListNode(first *Node) {
	if first.Next == nil {
		fmt.Println("empty node list")
		return
	}

	current := first
	for {
		fmt.Printf("node %d -> ", current.Id)
		if current.Next == first {
			break
		}
		current = current.Next
	}
	fmt.Println()
}

func CountNode(first *Node) int {
	count := 0

	if first.Next == nil {
		return count
	}

	current := first
	for {
		count++
		if current.Next == first {
			break
		}
		current = current.Next
	}
	return count
}

func Run(first *Node, start int, count int) {
	if first.Next == nil {
		fmt.Println("empty node list")
		return
	}

	if start > CountNode(first) {
		fmt.Println("start no cannot be greater than total no")
		return
	}

	tail := first
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}

	for i := 1; i <= start-1; i++ {
		first = first.Next
		tail = tail.Next
	}

	for {
		for i := 1; i <= count-1; i++ {
			first = first.Next
			tail = tail.Next
		}

		fmt.Printf("Node %d exit\n", first.Id)
		first = first.Next
		tail.Next = first

		if tail == first { // There is one node left.
			break
		}
	}

	fmt.Printf("Node %d exit\n", first.Id)
}

func main() {
	first := AddNode(100)
	ListNode(first)
	Run(first, 20, 51)
}
