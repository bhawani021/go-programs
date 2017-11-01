// Find length of linked list
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

// Get linked list length using iterative
func getLength(n *node) (length int) {
	for n != nil {
		length++
		n = n.Next
	}
	return
}

// Get linked list length using recursive
func getLength1(n *node) int {
	if n == nil {
		return 0
	}

	return 1 + getLength1(n.Next)
}

func main() {
	// Create linked list
	var n, current *node
	n = &node{Data: 0}
	current = n
	for i := 5; i <= 25; i += 5 {
		newNode := &node{Data: i}
		current.Next = newNode
		current = newNode
	}

	// Print linked list
	fmt.Println("Print linked list")
	displayNodes(n)

	// Print length of linked list
	length := getLength(n)
	fmt.Println("Length using iterative:", length)

	// Print lenth of linked list using recursive
	length = getLength1(n)
	fmt.Println("Length using recursive:", length)
}
