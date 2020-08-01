package main

import "fmt"

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

// Bubble sort
func sort(arr []int) {
	// Length of array
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}

func getPair(arr []int, sum int) (int, int) {
	// Length of array
	min, max := 0, len(arr)-1

	for min < max {
		if arr[min]+arr[max] == sum {
			return min, max
		}
		if arr[min]+arr[max] > sum {
			max--
		} else {
			min++
		}
	}
	return -1, -1
}

func main() {
	arr := []int{1, 2, 3, 4, 9, -3, 1, -9}

	sort(arr)

	fmt.Println(arr)

	i, j := getPair(arr, 10)

	fmt.Println(i, j)

}
