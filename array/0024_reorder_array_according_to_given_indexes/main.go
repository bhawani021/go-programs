package main

import "fmt"

// Reorder an array according to given indexes
func Rearrange(arr, index []int) {

	// Length of Array
	n := len(arr)

	for i := 0; i < n; i++ {

		for index[i] != i {

			ind := index[index[i]]
			a := arr[index[i]]

			index[index[i]] = index[i]
			arr[index[i]] = arr[i]

			index[i] = ind
			arr[i] = a
		}
	}

}


func main() {
	arr := []int{50, 40, 70, 60, 90}
	index := []int{3,  0,  4,  1,  2}

	Rearrange(arr, index)

	fmt.Println(arr) // [40 60 90 50 70]
	fmt.Println(index) // [0 1 2 3 4]
}

