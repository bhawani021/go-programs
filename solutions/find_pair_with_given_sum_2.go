package solutions

import "fmt"

func sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
}

func GetPair1(arr []int, sum int)  {
	// Sort given array
	sort(arr)
	fmt.Println("sorted arr:", arr)

	low := 0
	high := len(arr)-1

	for {
		if low < high {
			if arr[low] + arr[high] == sum {
				fmt.Printf("Indexes: %d and %d\n", low, high)
				return
			}
		} else {
			break
		}

		if arr[low] + arr[high] < sum{
			low++
		} else {
			high--
		}
	}

	fmt.Println("Could not found result")
}
