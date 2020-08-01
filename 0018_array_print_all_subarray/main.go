package main

import "fmt"

func printSubArray(arr []int) {
	// Length of array
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i; j <= n; j++ {
			for k := i; k < j; k++ {
				fmt.Print(arr[k], " ")
			}
			fmt.Println("")
		}
	}
}
func main() {
	arr := []int{1, 2, 3, 4}

	printSubArray(arr)
}
