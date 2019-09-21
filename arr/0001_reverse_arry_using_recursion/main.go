package main

import "fmt"

func reverse(arr []int, start, end int) {
	if start >= end {
		return
	}
	// Swap elements
	temp := arr[start]
	arr[start] = arr[end]
	arr[end] = temp

	reverse(arr, start+1, end-1)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	reverse(arr, 0, len(arr)-1)

	fmt.Println(arr)
}
