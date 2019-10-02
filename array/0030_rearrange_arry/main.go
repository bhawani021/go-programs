package main

import "fmt"

// Rearrange an array such that ‘arr[j]’ becomes ‘i’ if ‘arr[i]’ is ‘j’
func Rearrange(arr []int) {
	// Length of array
	n := len(arr)

	for i := 0; i < n; i++ {
		val := arr[i] % n
		arr[val] += i * n
	}

	for i := 0; i < n; i++ {
		arr[i] /= n
	}
}

func main() {
	arr := []int{2, 0, 1, 4, 5, 3}
	Rearrange(arr)
	fmt.Println(arr) // [1 2 0 5 3 4]
}
