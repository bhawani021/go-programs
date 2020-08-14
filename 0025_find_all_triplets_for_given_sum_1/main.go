package main

import "fmt"

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
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

func printTriplets(arr []int, targetSum int) {
	// Length of array
	n := len(arr)

	for i := 0; i < n-2; i++ {
		sum := targetSum - arr[i]
		start := i + 1
		end := n - 1
		for start < end {
			t := arr[start] + arr[end]
			if t == sum {
				fmt.Println("[", arr[start], arr[end], arr[i], "]")
				start++
				end--
			} else if t < sum {
				start++
			} else {
				end--
			}
		}
	}
	fmt.Println("Execution complete")
}

func main() {
	arr := []int{1, -1, 2, 3, -2, 4, 1, -3}
	fmt.Println(arr)
	sort(arr)
	printTriplets(arr, 2)
	/*
		[1 -1 2 3 -2 4 1 -3]
		[ 1 -1 2 ]
		[ 1 3 -2 ]
		[ 1 4 -3 ]
		[ -1 2 1 ]
		[ 2 3 -3 ]
		[ 3 -2 1 ]
		[ 4 1 -3 ]
	*/
}

// Complexity: O(n3)
