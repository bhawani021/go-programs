package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50, 60}

	start := 0
	end := len(arr) - 1

	for start < end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}

	fmt.Println(arr)
}
