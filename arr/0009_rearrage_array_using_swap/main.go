package main

import "fmt"

// Rearrange an array such that arr[i] = i
func ReArrayArrayUsingSwap(arr []int) {
	for i := 0; i < len(arr); {

		if arr[i] >= 0 && arr[i] != i {
			temp := arr[arr[i]]
			arr[arr[i]] = arr[i]
			arr[i] = temp
		} else {
			i++
		}
	}
}

func main()  {
	arr := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
	ReArrayArrayUsingSwap(arr)
	fmt.Println(arr) // [-1 1 2 3 4 -1 6 -1 -1 9]
}