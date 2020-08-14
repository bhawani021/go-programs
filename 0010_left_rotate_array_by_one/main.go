package main

import "fmt"

func leftRotate(arr []int, d int) {
	for i := 0; i < d; i++ {
		leftRotateByOne(arr)
	}
}

func leftRotateByOne(arr []int) {
	temp := arr[0]

	// Length of array
	n := len(arr)

	for i := 0; i < n-1; i++ {
		arr[i] = arr[i+1]
	}

	arr[n-1] = temp
	// or arr[i] = temp
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	leftRotate(arr, 2)
	fmt.Println(arr)
}
