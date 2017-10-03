package array

import (
	"fmt"
	"math/rand"
)

// SortArrayInDescendingOrder ...
func SortArrayInDescendingOrder() []int {
	var nums []int
	var n int
	fmt.Print("Please enter total numbers of elements in array: ")
	fmt.Scanf("%d", &n)

	// Insert elements into array
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Intn(100))
	}
	fmt.Println("Array before Sorting", nums)

	// Sort array in descending order
	var temp int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] < nums[j] {
				temp = nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
	return nums
}
