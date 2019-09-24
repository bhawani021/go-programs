package main

import "fmt"

func sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println("Sorted array:", arr)
}

func GetPair(arr []int, sum int) (int, int) {
	// Sort array
	sort(arr)
	min := 0
	max := len(arr)-1

	for min < max {
		if arr[min] + arr[max] == sum {
			return min, max
		}

		if arr[min] + arr[max] > sum {
			max--
		} else {
			min++
		}
	}
	return -1, -1
}

func main() {
	arr := []int{1, 0, 4, 6, 2}

	i, j := GetPair(arr, 10)
	fmt.Println(i, j)
}