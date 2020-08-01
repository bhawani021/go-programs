package main

import "fmt"

func getPair(arr []int, sum int) (int, int) {
	hash := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if val, ok := hash[sum-arr[i]]; ok {
			return val, i
		}
		hash[arr[i]] = i
	}

	return -1, -1
}
func main() {
	arr := []int{1, 2, 3, 0, 5, 9, 8}

	i, j := getPair(arr, 5)

	fmt.Println(i, j)
}
