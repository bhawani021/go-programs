package main

import "fmt"

func MoveAllZeroToEnd(arr *[5]int) {
	var count int // count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			temp := arr[count]
			arr[count] = arr[i]
			arr[i] = temp
			count++
		}
	}
}

func main() {
	var arr[5]int

	arr[0] = 0
	arr[1] = 3
	arr[2] = 1
	arr[3] = 0
	arr[4] = 5

	MoveAllZeroToEnd(&arr)

	fmt.Println(arr) // [3 1 5 0 0]
}
