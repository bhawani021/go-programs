package main

import "fmt"

// Node ...
type Node struct {
	value int
	next  *Node
}

func insertNodeAtEnd(n *Node, val int) {
	newNode := &Node{value: val}

	temp := n
	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode

	n = temp
}

func printList(n *Node) {
	for n != nil {
		fmt.Print(n.value, " ")
		n = n.next
	}
	fmt.Println("")
}

func main() {
	head := new(Node)
	head.value = 10

	second := &Node{value: 20}
	head.next = second

	third := &Node{value: 30}
	second.next = third

	printList(head)

	insertNodeAtEnd(head, 40)
	printList(head)
	insertNodeAtEnd(head, 50)
	printList(head)
}
