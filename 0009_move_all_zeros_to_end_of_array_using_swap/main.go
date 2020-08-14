package main

import "fmt"

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func rearrange(arr []int) {
	// Length of array
	n := len(arr)

	count := 0

	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			swap(arr, i, count)
			count++
		}
	}
}

func main() {
	arr := []int{0, 1, 9, 8, 4, 0, 0, 2, 7, 0, 6, 0, 9}
	rearrange(arr)
	fmt.Println(arr) // [1 9 8 4 2 7 6 9 0 0 0 0 0]
}

// Time Complexity: O(n).
// Auxiliary Space: O(1).
