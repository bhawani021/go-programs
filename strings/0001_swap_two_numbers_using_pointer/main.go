package main

import "fmt"

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	var a, b int
	a = 10
	b = 20
	swap(&a, &b)
	fmt.Println(a, b)
}
