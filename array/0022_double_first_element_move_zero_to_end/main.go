/* Double the first element and move zero to end

Given an array of integers of size n. Assume ‘0’ as invalid number and all other as valid number.
Convert the array in such a way that if next number is a valid number and same as current number, double its value and replace the next number with 0.
After the modification, rearrange the array such that all 0’s are shifted to the end.

 */
package main

import "fmt"

func moveZeroToEnd(arr []int) {
	var count int

	for i := 0; i < len(arr); i++ {

		if arr[i] != 0 {
			arr[count] = arr[i]
			count++
		}
	}

	for count < len(arr) {
		arr[count] = 0
		count++
	}

}

func ReArrangeArray(arr []int) {
	// Get length of Array
	n := len(arr)

	if n == 1 {
		return
	}

	for i := 0; i < n-1; i++ {

		if arr[i] != 0 && arr[i] == arr[i+1] {
			// Double current index value
			arr[i] = 2 * arr[i]
			// Move zero to next value
			arr[i+1] = 0
			// increment by 1 so as to move two indexes ahead during loop iteration
			i++
		}
	}

	fmt.Println(arr)

	moveZeroToEnd(arr)
}

func main()  {
	arr := []int{0, 2, 2, 2, 0, 6, 6, 0, 0, 8}
	ReArrangeArray(arr)

	fmt.Println(arr)
}
