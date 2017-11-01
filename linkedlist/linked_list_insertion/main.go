/*
Insert value in a linked list
1. At beginning
2. After a key
3. At end
*/

package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func displayNode(node *Node) {
	for node != nil {
		fmt.Print(node.Data, " ")
		node = node.Next
	}
	fmt.Println("")
}

func addNodeAtBegining(node *Node, key int) *Node {
	var head *Node
	head = &Node{Data: key, Next: nil}
	head.Next = node

	return head
}

func addNodeAfterNode(prev *Node, key int) {
	var newNode *Node
	newNode = &Node{Data: key, Next: nil}

	newNode.Next = prev.Next
	prev.Next = newNode
}

func addNodeAtEnd(node *Node, key int) *Node {
	var newNode *Node
	newNode = &Node{Data: key, Next: nil}

	temp := node
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = newNode

	return node
}

func main() {
	// Create four nodes
	var head, second, third, forth *Node
	head = &Node{Data: 10}
	second = &Node{Data: 30}
	third = &Node{Data: 40}
	forth = &Node{Data: 50}

	// Link first node to second
	head.Next = second
	// Link second node to third
	second.Next = third
	// Link third node to forth
	third.Next = forth

	// Print all nodes
	fmt.Println("All nodes before insertion.")
	displayNode(head)

	// Add node at beginning
	fmt.Println("Insert a new node at beginning.")
	head = addNodeAtBegining(head, 5)
	displayNode(head)

	// Add node after a node
	fmt.Println("Insert a node after a node.")
	addNodeAfterNode(second, 35)
	displayNode(head)

	// Add node at end
	fmt.Println("Add a new node at end.")
	head = addNodeAtEnd(head, 60)
	displayNode(head)
}
