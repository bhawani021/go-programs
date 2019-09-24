package main

import "fmt"

func ArrayCyclicallyRotation(arr []int) {
	// Length of array
	n := len(arr)
	// Take last value in temp
	temp := arr[n-1]

	for i := n-1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = temp
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	ArrayCyclicallyRotation(arr)
	fmt.Println(arr)
}