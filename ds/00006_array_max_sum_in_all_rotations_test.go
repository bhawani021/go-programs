// Get sum of array elements in format of arr[i] * i in all possible rotation
package ds

import "testing"

func TestMaxSumOfArrayInAllRotation(t *testing.T) {
	arr := []int{1, 2, 3}

	result := MaxSumOfArrayInAllRotation(arr)

	expected := 8

	if result != expected {
		t.Errorf("Result: %d, expected result: %d", result, expected)
	}
}
