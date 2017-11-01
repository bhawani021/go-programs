// Create binary tree
package main

// node structure
type node struct {
	Data  int
	Left  *node
	Right *node
}

func main() {
	// Create three nodes
	var root, left, right *node
	root = &node{Data: 1}
	left = &node{Data: 2}
	right = &node{Data: 3}

	// Link root node with left node
	root.Left = left
	// Link root node with right node
	root.Right = right
}
