// Delete a node from linked list using position
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
	fmt.Println(" ")
}

func deleteNode(n *node, position int) *node {
	if n == nil {
		return n
	}

	// If head need to be removed
	if position == 0 {
		n = n.Next
		return n
	}

	// Get previous node of a node to be removed
	var prev *node
	prev = n
	for i := 0; i < position-1; i++ {
		if prev == nil {
			break
		}
		prev = prev.Next
	}

	// If position is more than number of nodes.
	if prev == nil {
		return n
	}

	if prev.Next != nil {
		prev.Next = prev.Next.Next
	} else {
		// IF last element need to be removed.
		prev.Next = nil
	}

	return n
}

func main() {
	// Create linked list
	var n, current *node
	n = &node{Data: 0, Next: nil}
	current = n
	for i := 5; i <= 50; i += 5 {
		newNode := &node{Data: i, Next: nil}
		current.Next = newNode
		current = newNode
	}

	// Print linked list
	fmt.Println("Print linked list before deletion")
	displayNodes(n)

	// Delete head node
	fmt.Println("Delete head node")
	n = deleteNode(n, 0)
	displayNodes(n)

	fmt.Println("Delete an another node")
	n = deleteNode(n, 4)
	displayNodes(n)
}
