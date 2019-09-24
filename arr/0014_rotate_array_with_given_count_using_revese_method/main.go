package main

import "fmt"

func RotateArrayUsingReverseAlgorithm(arr []int, d int) {
	length := len(arr)
	reverseArray(arr, 0, d)
	reverseArray(arr, d, length)
	reverseArray(arr, 0, length)
}

func reverseArray(arr []int, start, end int) {
	fmt.Println(start, end)
	for {
		if start < end-1 {
			temp := arr[start]
			arr[start] = arr[end-1]
			arr[end-1] = temp
			start++
			end--
		} else {
			break
		}

	}
}


func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	RotateArrayUsingReverseAlgorithm(arr, 3)
	fmt.Println(arr)
}