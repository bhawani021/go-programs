// Delete a node from linked list using key

package main

import "fmt"

type node struct {
	Data int
	Next *node
}

func displayNodes(n *node) {
	for n != nil {
		fmt.Print(n.Data, " ")
		n = n.Next
	}
	fmt.Println("")
}

func deleteNode(n *node, key int) *node {
	if n == nil {
		return n
	}

	// If head node itself contains given key
	if n.Data == key {
		n = n.Next
		return n
	}

	current := n
	var prev *node
	for current != nil {
		if current.Data == key {
			break
		}
		prev = current
		current = current.Next

	}

	// If key was not present in linked list
	if current == nil {
		return n
	}

	prev.Next = current.Next

	return n
}

func main() {
	// Create a new linked list
	var n, current *node
	n = &node{Data: 0, Next: nil}
	current = n
	for i := 5; i <= 50; i += 5 {
		newNode := &node{Data: i, Next: nil}
		current.Next = newNode
		current = newNode
	}

	// Display linked list
	fmt.Println("Print linked list before delete")
	displayNodes(n)

	// Delete a node from linked list
	n = deleteNode(n, 40)
	// Print nodes after deletion
	fmt.Println("Print linked list after deleting a node")
	displayNodes(n)
}
