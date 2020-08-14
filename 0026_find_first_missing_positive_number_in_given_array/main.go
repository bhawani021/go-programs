package main

import "fmt"

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func findFirstMissingPostiveNumber(arr []int) int {
	// Length of array
	n := len(arr)
	if n == 0 {
		return -1
	}

	var i int
	for i < n {
		// Take element of array greater than 0
		// Take element of array if it's less than length of array + 1
		// Take element wich is not equal to i+1
		// [1, 2, 3, 4, 5, 6, 7] -> [i+1], ideal case
		if arr[i] > 0 && arr[i] < n+1 && arr[i] != arr[arr[i]-1] {
			swap(arr, i, arr[i]-1)
		} else {
			i++
		}

		fmt.Println(arr)
	}

	for i := 0; i < n; i++ {
		if arr[i] != i+1 {
			return i + 1
		}
	}

	return -1

}

func main() {
	arr := []int{-2, 11, 1, -3, 2, 10, 5}
	fmt.Println(arr)
	fmt.Println("==============================")
	num := findFirstMissingPostiveNumber(arr)
	fmt.Println(num)
}
/*
[-2 11 1 -3 2 10 5]
==============================
[-2 11 1 -3 2 10 5]
[-2 11 1 -3 2 10 5]
[1 11 -2 -3 2 10 5]
[1 11 -2 -3 2 10 5]
[1 11 -2 -3 2 10 5]
[1 2 -2 -3 11 10 5]
[1 2 -2 -3 11 10 5]
[1 2 -2 -3 11 10 5]
[1 2 -2 -3 5 10 11]
[1 2 -2 -3 5 10 11]
3
*/