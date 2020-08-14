package main

import "fmt"

func printTriplets(arr []int, sum int) {

	// Length of array
	n := len(arr)

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if arr[i]+arr[j]+arr[k] == sum {
					fmt.Println("[", arr[i], arr[j], arr[k], "]")
				}
			}
		}
	}
	fmt.Println("Execution complete")
}

func main() {
	arr := []int{1, -1, 2, 3, -2, 4, 1, -3}
	fmt.Println(arr)
	printTriplets(arr, 2)
	/*
		[1 -1 2 3 -2 4 1 -3]
		[ 1 -1 2 ]
		[ 1 3 -2 ]
		[ 1 4 -3 ]
		[ -1 2 1 ]
		[ 2 3 -3 ]
		[ 3 -2 1 ]
		[ 4 1 -3 ]
	*/
}

// Complexity: O(n3)
