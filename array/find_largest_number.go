package array

import (
	"fmt"
)

// FindLargestNumber ...
func FindLargestNumber(nums []int) (n int) {
	if len(nums) < 1 {
		fmt.Println("No element found.")
		return
	}

	if len(nums) == 1 {
		return nums[0]
	}

	n = nums[0]
	for _, v := range nums {
		if v > n {
			n = v
		}
	}
	return
}
