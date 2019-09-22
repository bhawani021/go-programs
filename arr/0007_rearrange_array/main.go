package main

import "fmt"

// Rearrange an array such that arr[i] = i
func RearrangeArray(arr []int) {
	// Length of array
	n := len(arr)
	for i := 0; i < n; i++ {

		if arr[i] != -1 && arr[i] != i {

			x := arr[i]

			for arr[x] != -1 && arr[x] != x {
				y := arr[x]

				arr[x] = x

				x = y
			}
			arr[x] = x

			// Check if not correct value for arr[i]
			if arr[i] != i {
				arr[i] = -1
			}
		}
	}
}

func main()  {
	arr := []int{0, -1, 2, 1}
	RearrangeArray(arr)
	fmt.Println(arr) // [0 1 2 -1]
}
