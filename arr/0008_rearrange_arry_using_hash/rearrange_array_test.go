package main

import (
	"reflect"
	"testing"
)

func TestRearrangeArray(t *testing.T) {
	inputArr := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4,-1}

	RearrangeArray(inputArr)

	expectedResult := []int{-1, 1, 2, 3, 4, -1, 6, -1, -1, 9}

	// Compare two array
	if !reflect.DeepEqual(inputArr, expectedResult) {
		t.Errorf("Expected: %v, found: %v", expectedResult, inputArr)
	}
}

