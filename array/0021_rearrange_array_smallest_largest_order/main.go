package main

import "fmt"

func swap(arr []int, first, second int) {
	temp := arr[first]
	arr[first] = arr[second]
	arr[second] = temp
}

func sort(arr []int) {
	// Length of array
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := i+1; j < n; j++ {
			if arr[i] > arr[j] {
				swap(arr, i, j)
			}
		}
	}
}

// Rearrange an array in order – smallest, largest, 2nd smallest, 2nd largest
func Rearrange(arr []int) []int {

	// Sort array
	sort(arr)

	// Length of array
	n := len(arr)
	result := make([]int, n)

	min := 0
	max := n-1
	index := 0
	for i := 0; i < n/2; i++ {
		result[index] = arr[min]
		min++
		index++
		result[index] = arr[max]
		max--
		index++
	}

	return result
}

func main() {
	arr := []int{0, 1, 9, 4, 8, 1}
	res := Rearrange(arr)
	fmt.Println(res) // [0 9 1 8 1 4]
}
