package main

import "fmt"

func getPair(arr []int, sum int) (int, int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] + arr[j] == sum {
				return i, j
			}
		}
	}
	return -1, -1
}

func main()  {
	arr := []int{1, 2, 3, 4, 5, 6}
	i, j := getPair(arr, 5)
	fmt.Println(i, j) // 0 3
}
