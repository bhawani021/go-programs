package main

import "fmt"

type node struct {
	value int
	next  *node
}

func printList(n *node) {
	for n != nil {
		fmt.Println(n.value)
		n = n.next
	}
}

func main() {
	head := &node{value: 10, next: nil}
	second := &node{value: 20, next: nil}
	third := &node{value: 30, next: nil}

	head.next = second
	second.next = third

	printList(head)
}
