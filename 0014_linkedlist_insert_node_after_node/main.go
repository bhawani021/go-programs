package main

import "fmt"

// Node ...
type Node struct {
	value int
	next  *Node
}

func printList(node *Node) {
	for node != nil {
		fmt.Print(node.value, " ")
		node = node.next
	}
	fmt.Println("")
}

func insertNodeAfterNode(prev *Node, val int) {
	newNode := &Node{value: val}

	// New node next will point to prev node next
	newNode.next = prev.next

	// Prev node next will point to new node
	prev.next = newNode
}

func main() {
	head := new(Node)
	head.value = 10

	second := &Node{value: 20}
	third := &Node{value: 30}
	fourth := &Node{value: 50}

	head.next = second
	second.next = third
	third.next = fourth

	insertNodeAfterNode(third, 40)
	printList(head)

}
