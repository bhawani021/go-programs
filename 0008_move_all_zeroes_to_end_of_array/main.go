// Move all zero at end of array
package main

import "fmt"

func rearrange(arr []int) {
	// Length of array
	n := len(arr)

	// Count of non-zero elements
	count := 0

	// Traverse the array. If element encountered is
	// non-zero, then replace the element at index 'count'
	// with this element
	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			arr[count] = arr[i]
			count++
		}
		fmt.Println(arr)
	}

	// Now all non-zero elements have been shifted to
	// front and 'count' is set as index of first 0.
	// Make all elements 0 from count to end.
	for count < n {
		arr[count] = 0
		count++
	}
}

func main() {
	arr := []int{1, 2, 0, 4, 3, 0, 5, 0}
	rearrange(arr)
	fmt.Println(arr) // [1 2 4 3 5 0 0 0]
}

// Time Complexity: O(n) where n is number of elements in input array.
// Auxiliary Space: O(1)
