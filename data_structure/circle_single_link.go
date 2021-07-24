package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		fmt.Println(newCatNode, " added to the link")
		return
	}

	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	temp.next = newCatNode
	newCatNode.next = head
	fmt.Println(newCatNode, " added to the link")
}

func ListCatNode(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("empty link")
		return
	}
	for {
		fmt.Printf("[Cat %d, %s] ->", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	fmt.Println()
}

func DeleteCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	if temp.next == nil { // empty node
		fmt.Println("empty link")
		return head
	}

	if temp.next == head { // has only one node
		temp.next = nil
		return head
	}

	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	deleted := false
	for {
		if temp.next == head {
			break
		}
		if temp.no == id {
			if temp == head {
				head = head.next
			}
			helper.next = temp.next
			fmt.Printf("deleted [cat %d, %s]\n", temp.no, temp.name)
			deleted = true
			break
		}
		temp = temp.next
		helper = helper.next
	}
	if !deleted {
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("deleted [cat %d, %s]\n", temp.no, temp.name)
		} else {
			fmt.Printf("cat %d not existing\n", id)
		}
	}

	return head
}

func main() {
	head := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "Tom",
	}

	cat2 := &CatNode{
		no:   2,
		name: "Tobby",
	}

	cat3 := &CatNode{
		no:   3,
		name: "Tim",
	}

	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCatNode(head)

	head = DeleteCatNode(head, 3)
	fmt.Println("head node=", head)
	ListCatNode(head)
}
