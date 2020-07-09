package ds

import (
	"reflect"
	"testing"
)

func TestReverseArrWithRecursion(t *testing.T) {
	arr := []int{10, 20, 30, 40, 50}

	reverseArrWithRecursion(arr, 0, len(arr)-1)

	expectedResult := []int{50, 40, 30, 20, 10}

	if !reflect.DeepEqual(arr, expectedResult) {
		t.Errorf("Result: %v, expected result: %v", arr, expectedResult)
	}
}
