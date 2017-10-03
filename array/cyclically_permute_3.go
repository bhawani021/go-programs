package array

import (
	"fmt"
	"math/rand"
)

// CyclicallyPermute ...
func CyclicallyPermute() [10]int {
	var nums [10]int
	var n int
	fmt.Print("Enter number of elements: ")
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(50)
	}
	fmt.Println("Before Cyclically Permute: ", nums)

	var temp int
	temp = nums[0]
	for i := 0; i < n; i++ {
		nums[i] = nums[i+1]
	}
	nums[n-1] = temp
	fmt.Println("After Cyclically Permute", nums)

	return nums
}
