package ds

// MaxSumOfArrayInAllRotation return sum of array elements in all possible rotation
// Get sum of array elements in format of arr[i] * i in all possible rotation
func MaxSumOfArrayInAllRotation(arr []int) int {
	// Length of array
	n := len(arr)
	var sum int
	for i := 0; i < n; i++ {

		var tempSum int
		for j := 0; j < n; j++ {
			index := (i + j) % n
			tempSum += arr[index] * j
		}

		if tempSum > sum {
			sum = tempSum
		}
	}
	return sum
}
