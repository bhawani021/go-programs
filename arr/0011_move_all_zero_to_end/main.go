package main

import "fmt"

func ArrayZeroToEnd(arr []int) {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[count] = arr[i]
			count++
		}
	}

	for count < len(arr) {
		arr[count] = 0
		count++
	}
}

func main() {
	arr := []int{1, 2, 2, 0, -1, 4, 9}
	ArrayZeroToEnd(arr)
	fmt.Println(arr)
}