package main

import "fmt"

// All negative integers appear before all the positive integers in the array
// without using any additional data structure like hash table, arrays.
// The order of appearance should be maintained.
func Rearrange(arr []int) {

	// Length of Array
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]

		if key > 0 {
			continue
		}

		j := i - 1

		for j >= 0 && arr[j] > 0 {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}

}

func main() {
	arr := []int{-12, 11, -13, -5, 6, -7, 5, -3, -6}

	Rearrange(arr)

	fmt.Println(arr)

}