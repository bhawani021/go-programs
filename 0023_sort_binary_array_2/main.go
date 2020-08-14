package main

import "fmt"

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func sort(arr []int) {
	// Length of array
	n := len(arr)

	pivot := 1
	var j int

	for i := 0; i < n; i++ {
		if arr[i] < pivot {
			swap(arr, i, j)
			j++
		}
	}
}

func main() {
	arr := []int{0, 1, 0, 0, 1, 1, 1, 1, 1, 0}

	sort(arr)

	fmt.Println(arr) // [0 0 0 0 1 1 1 1 1 1]
}

// Time complexity O(n)
// Space compexity O(1)
