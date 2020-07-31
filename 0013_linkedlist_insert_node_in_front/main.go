package main

import "fmt"

// Node ...
type Node struct {
	value 	int
	next    *Node
}

func insertNodeAtFront(n *Node, val int) *Node{
	head := &Node{value: val}
	head.next = n

	n = head

	return n
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
	
	second := &Node{value:20}
	head.next = second
	
	third := &Node{value:30}
	second.next = third

	printList(head) 

	n := insertNodeAtFront(head, 60)
	printList(n) 
}
