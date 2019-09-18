package main

import (
	"fmt"
	"go-programs/solutions"
)

func main() {
	arr := []int{1, 2, -3, 9, 0}
	res := solutions.ZeroSumSubArray(arr)
	fmt.Println(res) // true

	arr = []int{1, 2, 3, 4, 0}
	res = solutions.ZeroSumSubArray(arr)
	fmt.Println(res) // false
}