/* Maximum sum of i*arr[i] among all rotations of a given array

Input : arr = {1, 2, 3}
Output : 7

 */
package main

import "fmt"

func getMaxSumInAllRotation(arr []int) int {
	var res int
	// Length of array
	n := len(arr)
	for i := 0; i < n; i++ {
		var currentSum int

		for j := 0; j < n; j++ {
			index := (i + j) % n
			currentSum += arr[index] * j
		}

		if res < currentSum {
			res = currentSum
		}
	}

	return res
}

func main() {
	arr := []int{1, 2, 3}
	res := getMaxSumInAllRotation(arr)
	fmt.Println(res) // 7
}
