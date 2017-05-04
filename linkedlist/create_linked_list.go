package linkedlist

import "fmt"

// CreateLinkedList ...
func CreateLinkedList() *Node {
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

// Call
// var list = CreateLinkedList()
// printLinkedList(list)
