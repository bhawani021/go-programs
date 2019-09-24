package main

import "fmt"

func Search(arr []int) int {
	minIndex := -1
	min := arr[0]
	// Find index of min value in array
	for i := 0; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
			minIndex = i
		}
	}

	return minIndex
}

func main() {
	arr := []int{10, 1, 4, 5, 6}
	rotation := Search(arr)
	fmt.Println(rotation) // 1
}
