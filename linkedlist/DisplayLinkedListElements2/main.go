package main

import "fmt"

type node struct {
	value	int
	next    *node
}

func displayLinkedListElements() {
	repeat := 1
	var val int
	var head, current *node

	for repeat == 1 {
		fmt.Print("Please enter a node value: ")
		fmt.Scanf("%d", &val)

		// Create a new node
		var newNode *node
		newNode = &node{
			value: val,
			next: nil,
		}

		if head == nil {
			head = newNode
		} else {
			current.next = newNode
		}
		current = newNode

		fmt.Print("Do you want to continue? (Enter 1 to continue, 0 to exist): ")
		fmt.Scanf("%d", &repeat)
	}

	// Print list
	for n := head; n != nil; n = n.next {
		fmt.Printf("%d ", n.value)
	}
}

func main() {
	displayLinkedListElements()
}