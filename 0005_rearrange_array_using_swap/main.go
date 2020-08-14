/*
Input -> [-1, -1, 6, 1, 9, 3, 2, -1, 4, -1]
output -> [-1 1 2 3 4 -1 6 -1 -1 9]

[-1 -1 6 1 9 3 2 -1 4 -1]
[-1 -1 6 1 9 3 2 -1 4 -1]
[-1 -1 2 1 9 3 6 -1 4 -1]
[-1 -1 2 1 9 3 6 -1 4 -1]
[-1 1 2 -1 9 3 6 -1 4 -1]
[-1 1 2 -1 9 3 6 -1 4 -1]
[-1 1 2 -1 -1 3 6 -1 4 9]
[-1 1 2 -1 -1 3 6 -1 4 9]
[-1 1 2 3 -1 -1 6 -1 4 9]
[-1 1 2 3 -1 -1 6 -1 4 9]
[-1 1 2 3 -1 -1 6 -1 4 9]
[-1 1 2 3 -1 -1 6 -1 4 9]
[-1 1 2 3 4 -1 6 -1 -1 9]
[-1 1 2 3 4 -1 6 -1 -1 9]
[-1 1 2 3 4 -1 6 -1 -1 9]

1) Iterate through elements in array
2) If arr[i] >= 0 && arr[i] != i, put arr[i] at i ( swap arr[i] with arr[arr[i]])

*/
package main

import "fmt"

func main() {
	arr := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
	i, n := 0, len(arr)
	for i < n {

		if arr[i] >= 0 && arr[i] != i {
			temp := arr[arr[i]]
			arr[arr[i]] = arr[i]
			arr[i] = temp
		} else {
			i++
		}
		fmt.Println(arr)
	}
}
