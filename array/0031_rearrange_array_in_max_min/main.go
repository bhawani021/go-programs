package main

import "fmt"

/* Rearrange an array in maximum minimum form

Input: arr[] = {1, 2, 3, 4, 5, 6, 7}
Output: arr[] = {7, 1, 6, 2, 5, 3, 4}

*/

func Rearrange(arr []int) {
	// Length of array
	n := len(arr)

	low := 0
	high := n - 1

	// temp array
	tempArray := make([]int, len(arr))

	var flag bool
	flag = true

	for i := 0; i < n; i++ {
		if flag {
			tempArray[i] = arr[high]
			high--
		} else {
			tempArray[i] = arr[low]
			low++
		}

		flag = !flag
	}

	copy(arr, tempArray)
}

func main()  {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	Rearrange(arr)
	fmt.Println(arr) // [7, 1, 6, 2, 5, 3, 4]
}