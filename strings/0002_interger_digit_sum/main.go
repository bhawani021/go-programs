package main

import "fmt"

func sum(val int) int {
	sum := 0

	for val > 0 {
		sum += val%10
		val /= 10
	}

	return sum
}

func main() {
	val := 1234

	res := sum(val)

	fmt.Println(res) // 10
}