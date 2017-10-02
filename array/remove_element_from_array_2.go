package array

import (
	"fmt"
)

// RemoveElementFromArray ...
func RemoveElementFromArray() [10]int {
	var n int
	n = 5

	// Insert elements into array
	nums := [10]int{1, 8, 2, 7, 9}
	fmt.Println("Initial array", nums)

	// Sort array
	var temp int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[i] {
				temp = nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
	fmt.Println("Array after sorting", nums)

	// Element to be removed
	var key, found, pos int
	fmt.Print("Enter elements to remove: ")
	fmt.Scanf("%d", &key)
	for i := 0; i < n; i++ {
		if nums[i] == key {
			found = 1
			pos = i
			break
		}
	}

	if found == 1 {
		for i := pos; i < n-1; i++ {
			nums[i] = nums[i+1]
		}
	} else {
		fmt.Printf("Element %d is not found in array", key)
	}
	return nums
}
