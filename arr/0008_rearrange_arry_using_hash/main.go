package main

import "fmt"

// Rearrange an array such that arr[i] = i
func RearrangeArray(arr []int) {
	hash := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if _, exists := hash[arr[i]]; !exists {
			hash[arr[i]] = true
		}
	}
	fmt.Println("Hash: ", hash)

	for i := 0; i < len(arr); i++ {
		if _, exists := hash[i]; exists {
			arr[i] = i
		} else {
			arr[i] = -1
		}
	}
}

func main()  {
	arr := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
	RearrangeArray(arr)
	fmt.Println(arr) // [-1 1 2 3 4 -1 6 -1 -1 9]
}