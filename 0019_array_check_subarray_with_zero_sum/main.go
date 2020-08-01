package main

import "fmt"

func isZeroSumExists(arr []int) bool {
	hash := map[int]bool{0: true}

	var sum int

	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		if _, ok := hash[sum]; ok {
			return true
		}

		hash[sum] = true
	}
	return false
}

func main() {
	arr := []int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2}
	r := isZeroSumExists(arr)
	fmt.Println(r)

}
