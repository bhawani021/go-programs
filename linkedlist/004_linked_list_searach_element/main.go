// Search an element into linked list
package main

import "fmt"

type node struct {
	Data int
	Next *node
}

func display(n *node) {
	for n != nil {
		fmt.Print(n.Data, " ")
		n = n.Next
	}
	fmt.Println(" ")
}

// Search an element in linked list using iterative approach
func searchElement(n *node, key int) bool {
	if n == nil {
		return false
	}

	for n != nil {
		if n.Data == key {
			return true
		}
		n = n.Next
	}
	return false
}

// Search element in linked list using recursive approach
func searchElement1(n *node, key int) bool {

	if n == nil {
		return false
	}
	if n.Data == key {
		return true
	}

	return searchElement1(n.Next, key)
}

func main() {
	// Create a linked list
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
	display(n)

	// Search element using iterative approach
	v := 20
	res := searchElement(n, v)
	fmt.Printf("Search for element: %d using iterative, found? = %t\n", v, res)

	// Search element using reursive approach
	v = 21
	res = searchElement1(n, v)
	fmt.Printf("Search for element: %d using recursive, found? = %t\n", v, res)
}
