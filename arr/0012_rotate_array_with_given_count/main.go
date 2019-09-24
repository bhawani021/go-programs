package main

import "fmt"

func RotateArray(arr []int, count int) {
	for i := 0; i < count; i++ {
		RotateArrayByOne(arr)
	}
}

func RotateArrayByOne(arr []int) {
	temp := arr[0]
	var i int
	for i = 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[i] = temp
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	RotateArray(arr, 3)
	fmt.Println(arr)
}
