package array

import (
	"fmt"
)

// FilterEvenAndOddNumbers from a given array
func FilterEvenAndOddNumbers(nums [10]int) (a [10]int, b [10]int) {
	if len(nums) < 1 {
		fmt.Println("Elements not found in array.")
		return
	}

	var j, k int
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			a[j] = nums[i]
			j++
		} else {
			b[k] = nums[i]
			k++
		}
	}
	return
}
