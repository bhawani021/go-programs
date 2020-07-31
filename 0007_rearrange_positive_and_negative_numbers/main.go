/*
Rearrange positive and negative numbers
in O(n) time and O(1) extra space

An array contains both positive and negative numbers in random order.
Rearrange the array elements so that positive and negative numbers
are placed alternatively.
Number of positive and negative numbers need not be equal.
If there are more positive numbers they appear at the end of the
array. If there are more negative numbers, they too appear in the
end of the array.

For example, if the input array is [-1, 2, -3, 4, 5, 6, -7, 8, 9],
then the output should be [9, -7, 8, -3, 5, -1, 2, 4, 6]
*/
package main

import "fmt"

func swap(arr []int, first, second int) {
	temp := arr[first]
	arr[first] = arr[second]
	arr[second] = temp
}

func rearrange(arr []int) {
	// Length of array
	n := len(arr)
	i := -1

	for j := 0; j < n; j++ {
		if arr[j] < 0 {
			i++
			swap(arr, i, j)
		}
	}

	// Now all positive numbers are at end and negative numbers at
	// the beginning of array. Initialize indexes for starting point
	// of positive and negative numbers to be swapped
	neg := 0
	pos := i + 1

	// Increment the negative index by 2 and positive index by 1, i.e.,
	// swap every alternate negative number with next positive number
	for pos < n && neg < pos && arr[neg] < 0 {
		swap(arr, neg, pos)
		pos++
		neg += 2
	}
}

func main() {
	arr := []int{-1, 2, -3, 4, 5, 6, -7, 8, 9}
	rearrange(arr)
	fmt.Println(arr)
}
