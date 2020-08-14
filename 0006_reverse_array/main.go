package main

import "fmt"

func swap(arr []int, first, second int) {
	temp := arr[first]
	arr[first] = arr[second]
	arr[second] = temp
}

func reverseUsingIteration(arr []int) {
	min, max := 0, len(arr)-1

	for min < max {
		swap(arr, min, max)
		min++
		max--
	}
}

func reverseUsingRecursion(arr []int, min, max int) {
	if min > max {
		return
	}

	swap(arr, min, max)

	reverseUsingRecursion(arr, min+1, max-1)
}

func main() {
	// Iteration
	arr := []int{1, 2, 3, 4, 5}
	reverseUsingIteration(arr)
	fmt.Println(arr) // [5 4 3 2 1]

	// Recursion
	arr1 := []int{10, 20, 30, 40, 50}
	reverseUsingRecursion(arr1, 0, len(arr)-1)
	fmt.Println(arr1) // [50 40 3 20 10]
}

// Time Complexity : O(n)
