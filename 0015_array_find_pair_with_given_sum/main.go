package main

import "fmt"

func getPair(arr []int, sum int) (int, int) {
	// Length of array
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i]+arr[j] == sum {
				return i, j
			}
		}
	}
	return -1, -1
}

func main() {
	arr := []int{1, 2, 3, 4, 9}

	i, j := getPair(arr, 7)

	fmt.Println(i, j)

}
