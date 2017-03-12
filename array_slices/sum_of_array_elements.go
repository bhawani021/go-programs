// Sum of array elements

package main

import "fmt"

func main() {
	// Get sum of array elements using foo loop
	// Declare an array of interger tyoe
	nums := [5]int{1, 2, 3, 4, 5}
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	// Print sum
	fmt.Println("Sum =", sum)

	// Get sum of array elements using range function
	var total int
	for _, val := range nums {
		total += val
	}

	// Print total
	fmt.Println("Total =", total)
}