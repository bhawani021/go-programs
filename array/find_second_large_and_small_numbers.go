package array

import (
	"fmt"
)

// FindSecondLargestAndSmallestNumbers find second largest and second smallest number.
func FindSecondLargestAndSmallestNumbers(nums []int) (l int, s int) {
	if len(nums) < 4 {
		fmt.Println("There should be minimum 4 elements in array.")
		return
	}

	// Sort elements in ascending order.
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
	fmt.Println("After sorting: ", nums)

	// Get second smallest value.
	s = nums[1]
	// Get second largest value.
	l = nums[len(nums)-2]

	return l, s
}
