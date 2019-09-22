package main

import (
	"reflect"
	"testing"
)

func TestReArrayArrayUsingSwap(t *testing.T) {
	inputArr := []int{0, -1, 1}
	ReArrayArrayUsingSwap(inputArr)
	expectedResult := []int{0, 1, -1}

	if !reflect.DeepEqual(inputArr, expectedResult) {
		t.Errorf("Expected result: %v, found: %v", expectedResult, inputArr)
	}
}
