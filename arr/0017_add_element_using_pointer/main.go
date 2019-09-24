package main

import "fmt"

func Sum(arr *[5]int) int {
	var sum int

	for _, v := range arr {
		sum += v
	}

	return sum
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	res := Sum(&arr)

	fmt.Println(res)
}