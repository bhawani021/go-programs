package main

import "fmt"

// Rearrange an array such that ‘arr[j]’ becomes ‘i’ if ‘arr[i]’ is ‘j’
func Rearrange(arr []int) {
	// Length of array
	n := len(arr)
	tempArray := make([]int, n)

	for i := 0; i < n; i++ {
		tempArray[i] = arr[i]
	}

	for i := 0; i < n; i++ {
		// i = 0 arr[i] = 0
		arr[tempArray[i]] = i
	}
}

func main()  {
	arr := []int{3, 2, 1, 0}

	// Arr[0] = 1 -> Arr[1] = 0

	Rearrange(arr)

	fmt.Println(arr)
}
