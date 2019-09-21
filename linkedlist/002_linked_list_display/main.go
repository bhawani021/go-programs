package main

import "fmt"


type Node struct{
	Data int
	Next *Node
}


func main() {

	// Create three nodes
	var head, second, third *Node
	head = &Node{Data: 1}
	second = &Node{Data: 2}
	third = &Node{Data: 3}

	// Link first node with second
	head.Next = second 
	// Link second node with third
	second.Next = third

	// Print Linked list
	node := head
	for node != nil {
		fmt.Print(node.Data, " ")
		node = node.Next	
	}
	fmt.Println("")
}