package main

import "fmt"

// rearrange the array elements so that positive and
// negative numbers are placed alternatively.
func Rearrange(arr []int) {

	j := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			j++
			fmt.Println(i, j)
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	// Initial negative number index
	negativeVal := 0
	// Initial positive number index
	positiveVal := j + 1

	for negativeVal < positiveVal && positiveVal < len(arr) && arr[negativeVal] < 0 {
		temp := arr[negativeVal]
		arr[negativeVal] = arr[positiveVal]
		arr[positiveVal] = temp
		negativeVal += 2
		positiveVal++
	}
}

func main() {
	arr := []int{-1, 2, -3, 4, 5, 6, -7, 8, 9}
	Rearrange(arr)
	fmt.Println(arr)
}
