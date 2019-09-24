package main

import "fmt"

func GetPairWithSumUsingHash(arr []int, sum int) (int, int) {
	hash := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if _, exists := hash[sum-arr[i]]; exists {
			return hash[sum-arr[i]], i
		} else {
			hash[arr[i]] = i
		}
	}

	return -1, -1
}

func main()  {
	arr := []int{1, 2, 3, 5, 6, 7, 9}
	i, j := GetPairWithSumUsingHash(arr, 13)
	fmt.Println(i, j)
}
