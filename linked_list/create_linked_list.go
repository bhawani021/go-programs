package main

import "fmt"

// Node structure
type Node struct {
	Value int
	Next *Node
}

// Create linked list
func createLinkedList() *Node {
	var head = &Node{Value: 0, Next: nil}
	current := head

	for i := 5; i <= 100; i += 5 {
		newNode := &Node{Value: i, Next: nil}

		current.Next = newNode
		current = newNode
	}
	return head
}

// Print linked list
func printLinkedList(node *Node) {
	for current := node; current != nil; current = current.Next {
		fmt.Printf("%d ", current.Value)
	}
}

func main() {
	var list = createLinkedList()
	printLinkedList(list)
}