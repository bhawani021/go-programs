package main

import "fmt"

func sort(arr []int) {
	// Length of array
	n := len(arr)
	var zeros int

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			zeros++
		}
	}

	// Put zero at beginning of array
	for i := 0; i < zeros; i++ {
		arr[i] = 0
	}

	// Put one at end of array
	for i := zeros; i < n; i++ {
		arr[i] = 1
	}
}

func main() {
	arr := []int{0, 1, 0, 0, 1, 1, 1, 1, 1, 0}

	sort(arr)

	fmt.Println(arr) // [0 0 0 0 1 1 1 1 1 1]
}
