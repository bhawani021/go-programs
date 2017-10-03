package array

import (
	"fmt"
	"math/rand"
)

// SortArrayInAscendingOrder ...
func SortArrayInAscendingOrder() []int {
	var nums []int
	var n int
	fmt.Print("Please enter number of elements in array: ")
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Intn(100))
	}

	// Sort in ascending order
	var temp int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				temp = nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}

	return nums
}
