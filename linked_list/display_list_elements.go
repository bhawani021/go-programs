// Create a single linked list and display elements of list
package main

import "fmt"

// Node structure
type Node struct {
	Value int
	Next *Node
}

func main() {
	var choice, value int
	var head, current *Node

	choice = 1
	
	// Create linked list
	for choice == 1 {
		fmt.Print("Plase enter a number: ")
		fmt.Scanf("%d\n", &value)

		var newNode = &Node{value, nil}
		if head != nil {
			current.Next = newNode
		} else {
			head = newNode
		}
		current = newNode

		fmt.Print("Do you want to continue (1 for yes, 0 for no): ")
		fmt.Scanf("%d\n", &choice)	
	}

	// Print elements of linked list
	for node := head; node != nil; node = node.Next {
		fmt.Printf("%d ", node.Value)
	}
}