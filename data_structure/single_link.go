package main

import "fmt"

type HeroNode struct {
	no   int
	name string
	next *HeroNode
}

// Append nodes to the last
func AppendHeroNode(head, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	temp.next = newHeroNode
}

// Add nodes by order
func InsertHeroNode(head, newHeroNode *HeroNode) {
	temp := head
	canInsert := true

	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			break
		} else if temp.next.no == newHeroNode.no {
			canInsert = false
			break
		}
		temp = temp.next
	}

	if !canInsert {
		fmt.Println("cannot insert; ID exists already, no=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

func DeleteHeroNode(head *HeroNode, id int) {
	temp := head
	found := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			found = true
			break
		}
		temp = temp.next
	}

	if !found {
		fmt.Println("cannot find ID=", id)
		return
	} else {
		temp.next = temp.next.next
	}
}

func ListHeroNode(head *HeroNode) {
	temp := head

	if temp.next == nil {
		fmt.Println("empty link")
		return
	}

	for {
		fmt.Printf("[%d, %s] -> ", temp.next.no, temp.next.name)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
	fmt.Println()
}

func main() {
	head := &HeroNode{}

	hero1 := &HeroNode{
		no:   1,
		name: "Batman",
	}

	hero2 := &HeroNode{
		no:   2,
		name: "Superman",
	}

	hero3 := &HeroNode{
		no:   3,
		name: "Spiderman",
	}

	hero4 := &HeroNode{
		no:   3,
		name: "Ironman",
	}

	// AppendHeroNode(head, hero1)
	// AppendHeroNode(head, hero2)
	// AppendHeroNode(head, hero3)

	InsertHeroNode(head, hero3)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero4)
	InsertHeroNode(head, hero1)
	ListHeroNode(head)

	DeleteHeroNode(head, 2)
	DeleteHeroNode(head, 5)
	ListHeroNode(head)
}
