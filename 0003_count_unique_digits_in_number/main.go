package main

import "fmt"

// Input: N = 22342
// Output: 2
// Explanation:
// The digits 3 and 4 occurs only once. Hence, the output is 2.

// Input: N = 99677
// Output: 1
// Explanation:
// The digit 6 occurs only once. Hence, the output is 1.
//
//
//
// Input: 12345300
// Debug
// [1 0 0 0 0 0 0 0 0 0]
// [2 0 0 0 0 0 0 0 0 0]
// [2 0 0 1 0 0 0 0 0 0]
// [2 0 0 1 0 1 0 0 0 0]
// [2 0 0 1 1 1 0 0 0 0]
// [2 0 0 2 1 1 0 0 0 0]
// [2 0 1 2 1 1 0 0 0 0]
// [2 1 1 2 1 1 0 0 0 0]
// Output: 4

func getUniqueNumbers(num int) int {
	// An Array from 0-9 indexing with size 10
	var arr [10]int

	for num > 0 {
		digit := num % 10
		arr[digit]++
		fmt.Println(arr)
		num /= 10
	}

	var uniqueVal int
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			uniqueVal++
		}
	}
	return uniqueVal
}

func main() {
	num := getUniqueNumbers(12345300)
	fmt.Println(num)
}

// Time Complexity: O(N)
// Auxiliary Space: O(1)
