package main

import (
	"fmt"
)

func padding(n int) int {
	p := 1
	for p < n {
		p *= 10
	}
	return p
}


func compare(arr []int, i, j int) {
	val1 := arr[i]
	val2 := arr[j]
	c1 := val1 * padding(val2) + val2
	c2 := val2 * padding(val1) + val1
	if c2 > c1 {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}
}

func sort(arr []int) {
	// Length of array
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			compare(arr, i, j)
		}
	}
}

func main() {
	arr := []int{54, 546, 548, 60}
	fmt.Println(arr)

	sort(arr)

	fmt.Println(arr)
}


