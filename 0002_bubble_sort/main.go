package main

import "fmt"

/*
arr = [1, 2, 4, 0, 9, 7, 3]

Step  1
[1 2 4 0 9 7 3]
[1 2 4 0 9 7 3]
[1 2 0 4 9 7 3]
[1 2 0 4 9 7 3]
[1 2 0 4 7 9 3]
[1 2 0 4 7 3 9]

Step  2
[1 2 0 4 7 3 9]
[1 0 2 4 7 3 9]
[1 0 2 4 7 3 9]
[1 0 2 4 7 3 9]
[1 0 2 4 3 7 9]

Step  3
[0 1 2 4 3 7 9]
[0 1 2 4 3 7 9]
[0 1 2 4 3 7 9]
[0 1 2 3 4 7 9]

Step  4
[0 1 2 3 4 7 9]
[0 1 2 3 4 7 9]
[0 1 2 3 4 7 9]

Step  5
[0 1 2 3 4 7 9]
[0 1 2 3 4 7 9]

Step  6
[0 1 2 3 4 7 9]

Step  7
[0 1 2 3 4 7 9]
*/

// BubbleSort ...
func BubbleSort(arr []int) {
	// Length of array
	n := len(arr)
	if n == 0 || n == 1 {
		return
	}

	for i := 0; i < n; i++ {
		fmt.Println("Step ", i+1)
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
			fmt.Println(arr)
		}

	}
}

func main() {
	arr := []int{1, 2, 4, 0, 9, 7, 3}

	BubbleSort(arr)

	fmt.Println(arr)
}

// Complexity: O(n2)
