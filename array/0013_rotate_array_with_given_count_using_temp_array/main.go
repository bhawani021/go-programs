package main

import "fmt"

func RotateArray(arr []int, count int) []int {
	// Length of array
	n := len(arr)
	return append(arr[count:n], arr[0:count]...)
}

func main()  {
	arr := []int{1, 2, 3, 4, 5, 6}
	res := RotateArray(arr, 2)
	fmt.Println(res) // [3 4 5 6 1 2]
}