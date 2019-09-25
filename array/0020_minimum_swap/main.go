package main

import (
	"fmt"
)

// Minimum swaps required to bring all elements less than or equal to k together
func MinSwap(arr []int, k int) int {
	// Length of array
	n := len(arr)

	// Find elements which are less than or equal to k (good) and greater than or equal to k (bad)
	var good, bad int
	for i := 0; i < n; i++ {
		if arr[i] <= k{
			good++
		} else {
			bad++
		}
	}

	ans := bad
	j := good
	for i := 0; i < n; i++ {
		if j == n {
			break
		}

		if arr[i] > k {
			bad--
		}

		if arr[j] > k {
			bad++
		}

		if bad < ans {
			ans = bad
		}

		j++
	}

	return ans


}

func main() {
	arr := []int{2, 1, 5, 6, 3}

	r := MinSwap(arr, 3)

	fmt.Println(r)

}
