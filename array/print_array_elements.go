package array

import "fmt"

// PrintArrayElements ...
func PrintArrayElements() {
	// Declare a interger type array
	var nums [5]int

	// Assign value of array
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	nums[3] = 40
	nums[4] = 50

	// Print length of array
	fmt.Println("Length =", len(nums))

	// Print elements of array using in-built range function
	for i, v := range nums {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	// Print elements of array using for loop
	for i := 0; i < len(nums); i++ {
		fmt.Printf("index: %d, value: %d\n", i, nums[i])
	}
}
