package main

import "fmt"

// BinarySearch ...
func BinarySearch(arr []int, val, start, end int) int {
	if len(arr) == 0 || start > end {
		return -1
	}

	mid := (start + end) / 2

	if arr[mid] == val {
		return mid
	}

	if val > arr[mid] {
		return BinarySearch(arr, val, mid+1, end)
	}

	return BinarySearch(arr, val, start, mid-1)

}

func main() {
	arr := []int{10, 20, 30, 50, 60, 70}

	index := BinarySearch(arr, 50, 0, len(arr)-1)

	fmt.Println(index)
}

// Complexity: log(n)
// Space complexity: O(1)
