package main

import (
	"reflect"
	"testing"
)

func TestRearrange(t *testing.T) {
	inputArr := []int{1, -9, 2, 4, -7, 2}
	Rearrange(inputArr)
	expectedResult := []int{2, -7, 4, -9, 1, 2}

	if !reflect.DeepEqual(inputArr, expectedResult) {
		t.Errorf("Expected result: %v. Found: %v", expectedResult, inputArr)
	}
}