// Sum of array element using pointer as function arg
package main

import "fmt"

func main() {
	// Declare a interger type array
	nums := [5]int{10, 20, 30, 40}
	var result int

	result = sum(&nums)
	fmt.Println("Result =", result)
}

func sum(nums *[5]int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}