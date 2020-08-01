package main

import "fmt"

func sort(arr []int) {
	// Length of array
	n := len(arr)

	var k int

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			arr[k] = 0
			k++
		}
	}

	for i := k; i < n; i++ {
		arr[i] = 1
	}
}

func main() {
	arr := []int{0, 1, 0, 0, 1, 1, 1, 1, 1, 0}

	sort(arr)

	fmt.Println(arr) // [0 0 0 0 1 1 1 1 1 1]
}

// Time complexity O(n)
// Space compexity O(1)
