/*
Rearrange an array such that arr[i] = i using Hash

Input : arr = {-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
Output : [-1, 1, 2, 3, 4, -1, 6, -1, -1, 9]

Input : arr = {19, 7, 0, 3, 18, 15, 12, 6, 1, 8,
              11, 10, 9, 5, 13, 16, 2, 14, 17, 4}
Output : [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
         11, 12, 13, 14, 15, 16, 17, 18, 19]
*/
package main

import "fmt"

// [-1, -1, 6, 1, 9, 3, 2, -1, 4, -1]

func mapArrayElements(arr []int) map[int]bool {
	hash := make(map[int]bool)
	// Length of array
	n := len(arr)
	for i := 0; i < n; i++ {
		if _, ok := hash[arr[i]]; !ok {
			hash[arr[i]] = true
		}
	}

	return hash
}

func rearrangeArray(size int, hash map[int]bool) []int {
	var arr []int
	for i := 0; i < size; i++ {
		var val int
		if _, ok := hash[i]; ok {
			val = i
		} else {
			val = -1

		}
		arr = append(arr, val)
	}

	return arr
}

func main() {
	arr := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
	hash := mapArrayElements(arr)
	fmt.Println(hash)
	result := rearrangeArray(len(arr), hash)
	fmt.Println(result)
}
