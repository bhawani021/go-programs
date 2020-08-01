package main

import "fmt"

func printSubArrayWithSeroSum(arr []int) {
	// Length of array
	n := len(arr)
	for i := 0; i < n; i++ {
		var sum int
		for j := i; j < n; j++ {
			sum += arr[j]
			if sum == 0 {
				fmt.Printf("%d...%d\n", i, j)
			}
		}
	}
}

func main() {
	arr := []int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2}

	printSubArrayWithSeroSum(arr)
}
