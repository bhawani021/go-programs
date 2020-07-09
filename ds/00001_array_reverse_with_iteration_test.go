package ds

import (
	"reflect"
	"testing"
)

func TestReverseArrWithIteration(t *testing.T) {
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	reverseArrWithIteration(arr)

	expectedResult := []int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10}

	if !reflect.DeepEqual(arr, expectedResult) {
		t.Errorf("Result: %v, expected result: %v", arr, expectedResult)
	}
}
