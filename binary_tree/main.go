package main

import "fmt"

type Node struct {
	Id    int
	Name  string
	Left  *Node
	Right *Node
}

// PreOrder Output: root, left subtree, right subtree
func PreOrder(node *Node) {
	if node != nil {
		fmt.Printf("id=%d name=%s\n", node.Id, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

// InfixOrder Output: left subtree, root, right subtree
func InfixOrder(node *Node) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("id=%d name=%s\n", node.Id, node.Name)
		InfixOrder(node.Right)
	}
}

// PostOrder Output: left subtree, root, right subtree
func PostOrder(node *Node) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("id=%d name=%s\n", node.Id, node.Name)
	}
}

func main() {
	/*
							root node
							/		\
				left1 node			right1 node
				/		\					\
		left10 node		left11 node			right2 node
	*/
	root := &Node{
		Id:   1,
		Name: "Superman",
	}
	left1 := &Node{
		Id:   2,
		Name: "Batman",
	}

	left10 := &Node{
		Id:   10,
		Name: "Batman",
	}
	left11 := &Node{
		Id:   11,
		Name: "Batman",
	}
	left1.Left = left10
	left1.Right = left11

	right1 := &Node{
		Id:   3,
		Name: "Ironman",
	}
	root.Left = left1
	root.Right = right1

	right2 := &Node{
		Id:   4,
		Name: "Spiderman",
	}
	right1.Right = right2

	fmt.Println("******Preorder output:******")
	PreOrder(root)
	fmt.Println("******InfixOrder output:******")
	InfixOrder(root)
	fmt.Println("******PostOrder output:******")
	PostOrder(root)
}
