package main

import "fmt"

// Reorder an array according to given indexes
func Rearrange(arr, index []int) {

	tempArr := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		tempArr[index[i]] = arr[i]

	}

	for i := 0; i < len(arr); i++ {
		arr[i] = tempArr[i]
		index[i] = i
	}


}

func main() {
	arr := []int{50, 40, 70, 60, 90}
	index := []int{3,  0,  4,  1,  2}

	Rearrange(arr, index)

	fmt.Println(arr) // [40 60 90 50 70]
	fmt.Println(index) // [0 1 2 3 4]
}

