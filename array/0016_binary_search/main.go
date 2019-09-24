package main

import "fmt"

func BinarySearch(arr []int, needle int) int {
	// Length of array
	n := len(arr)

	min := 0;
	max := n

	for {
		if min < max {
			mid := (min + max) / 2

			if arr[mid] == needle {
				return mid
			}

			if arr[mid] < needle {
				min = mid + 1
			} else {
				max = mid
			}
		} else {
			break
		}
	}

	return -1
}

func main()  {
	data := []int{1, 2, 4, 5, 6}
	i := BinarySearch(data, 4)
	fmt.Println(i)
}
