package main

import "fmt"

type node struct {
	value int
	next  *node
}

func createLinkedList() *node {
	// Head of node
	head := &node{value:0, next:nil}

	current := head

	for i := 5; i <= 100; i += 5 {
		newNode := &node{
			value: i,
			next:nil,
		}

		current.next = newNode
		current = newNode
	}
	return head
}

func printLinkedList(n *node) {
	for current := n; current != nil; current = current.next {
		fmt.Printf("%d ", current.value)
	}
}

func main() {
	var n *node
	n = createLinkedList()
	printLinkedList(n)
}

