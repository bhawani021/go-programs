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

	for i := 0; i < n-1; i++ {
		for j := i+1; j < n; j++ {
			if arr[i] > arr[j] {
				swap(arr, i, j)
			}
		}
	}
}

// Rearrange array such that even positioned are greater than odd
func RearrangeArray(arr []int) []int {
	// Sort Array
	sort(arr)
	fmt.Println(arr)

	// Length of array
	n := len(arr)

	res := make([]int, n)

	p := 0
	q := n - 1
	for i := 0; i < n; i++ {
		// Assign even indexes with maximum elements
		if (i + 1) % 2 == 0 {
			res[i] = arr[q]
			q = q-1
		} else {
			res[i] = arr[p]
			p = p + 1
		}
	}

	return res
}

func main() {
	arr := []int{1, 3, 2, 2, 5}
	res := RearrangeArray(arr)
	fmt.Println(res)
}