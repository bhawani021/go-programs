package array

import (
	"fmt"
	"math/rand"
)

// InsertElementIntoArray ...
func InsertElementIntoArray() [10]int {
	var nums [10]int
	var n int
	fmt.Print("Please enter a number:")
	fmt.Scanf("%d", &n)

	// Insert elements into array
	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(100)
	}
	fmt.Println("Initial array:", nums)

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
	fmt.Println("After sorting:", nums)

	// Get a key
	var key int
	fmt.Print("Enter the element to be intested: ")
	fmt.Scanf("%d", &key)
	// Find correct position of key
	var pos int
	for i := 0; i < n; i++ {
		if key < nums[i] {
			pos = i
			break
		}
		if key > nums[n-1] {
			pos = n
			break
		}
	}

	var m int
	if pos != n {
		m = n - pos + 1
		for i := 0; i <= m; i++ {
			nums[n-i+2] = nums[n-i+1]
		}
	}
	nums[pos] = key
	return nums
}
